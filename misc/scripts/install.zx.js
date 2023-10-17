async function git(dir, repo, ref, files = []) {
  console.log(`pulling ${files.join(" ")} from ${repo} into ${dir}`);
  await $`mkdir -p ${dir}`;
  await within(async () => {
    cd(dir);
    await $`git init`;
    if (files.length > 0) {
      await $`git config core.sparsecheckout true`;
      await $`touch .git/info/sparse-checkout`;
      for (const file of files) {
        await $`echo "${file}" >> .git/info/sparse-checkout; rm -rf "${dir}/${file}"; echo "sparse ${file}"`;
      }
    }
    // ignore go test files
    $`echo "!**/*_test.go" >> .git/info/sparse-checkout;`
    await $`git remote add origin https://github.com/${repo}.git`;
    await $`git pull -s theirs --depth 1 origin ${ref}`;
    if (files.length > 0) {
      // only leave .git folder if we cloned the entire repo
      await $`rm -rf .git`;
    }
  });
}

async function cloneAwsSdkV1(ref) {
  // 6b43f9d073a68f8f4c5e5f6c273955730fe6bda4
  await $`rm -rf aws/aws-sdk-go`;
  await git("aws/aws-sdk-go", "aws/aws-sdk-go", ref, [
    // "go.mod",
    // "models",
    // "Gopkg.lock",
    // "aws",
    // "doc.go",
    // "go.sum",
    // "private",
    // "Gopkg.toml",
    // "awstesting",
    // // "example",
    // "internal",
    // "service/ec2",
    // "service/generate.go",
    // "service/iam",
    // "service/s3",
    // "service/sso",
    // "service/ssooidc",
    // "service/sts",
  ]);
}

// v1 and v2 because why not?
async function cloneAwsSdkV2(ref) {
  await $`rm -rf aws/aws-sdk-go-v2`;
  await git("aws/aws-sdk-go-v2", "aws/aws-sdk-go-v2", ref, [
    // "codegen",
    // "feature",
    // "config",
    // "go.mod",
    // "aws",
    // "credentials",
    // "go.sum",
    // "doc.go",
    // "internal",
    // "!service",
    // "service/ec2",
    // "service/s3",
  ]);
}

/**
 * Apply git patches from patchBaseDir directory. Searches patchBaseDir recursively
 * looking for files. All files are treated as git patch files. The subfolder location
 * in patchBaseDir is used as the base directory to apply the patches in. This relative
 * location (toPatch) from the project root should contain a `.git` folder.
 *
 * Do create new patch files cd into the folder to create the patch from and run `git diff > $PATCH_FILE`
 * Move the patch file into the .patches folder with a subdirectory that matches the relative path
 * of the cd directory
 *
 * @param {string} patchBaseDir Base directory to search for patches
 * @param {string} toPatch subfolder to search/apply patches in
 */
async function applyPatches(patchBaseDir, toPatch = "") {
  const dirents = await fs.readdir(path.join(patchBaseDir, toPatch), { withFileTypes: true });
  for (const dirent of dirents) {
    if (dirent.isDirectory()) {
      applyPatches(patchBaseDir, path.join(toPatch, dirent.name));
      continue;
    }
    within(async () => {
      cd(toPatch);
      await $`git apply -p1 ${path.join(patchBaseDir, toPatch, dirent.name)}`;
    });
  }
}

async function finalizeClone() {
  // TODO: this isn't working. zx sanitizing it?
  await $`rm -rf aws/**/*_test.go`;

  await applyPatches(path.resolve(__dirname, "..", "..", ".patches"));
}


async function install() {
  await Promise.all([
    cloneAwsSdkV1("v1.45.25"),
    // cloneAwsSdkV2("4599f78694cabb6853addabc6f92cb197fdb5647")
  ]);
  await finalizeClone()
  await $`go mod tidy`
  await $`ln -fs $PWD/src/mock.go $PWD/aws/aws-sdk-go/aws/request/mock.go`
}

install()
  .then(() => {
    console.log(chalk.bgGreen("install done"));
  })
  .catch((e) => {
    console.log(chalk.bgRed("install failed"));
    console.error(e);
    process.exit(1);
  });
