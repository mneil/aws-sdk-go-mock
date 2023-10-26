
async function patchWasmExec(outDir) {
  const goroot = await $`go env GOROOT`;
  const wasmExec = path.join(goroot.stdout.replace("\n", ""), "misc", "wasm", "wasm_exec.js");
  if (!wasmExec.endsWith("wasm_exec.js")) {
    throw new Error(
      "You must pass the absolute path to wasm_exec as the last argument to this script.\nnode ./scripts/wasm-exec.js $(go env GOROOT)/misc/wasm/wasm_exec.js"
    );
  }

  async function copyWasmExec(wasmExec, outDir) {
    // const globalName = "awsSdkGoMock";
    const providerName = "AwsSdkGoMock";
    let contents = await fs.promises.readFile(wasmExec, "utf-8");
    contents = contents.replace("(() => {", `export function ${providerName} () {`);

    // contents = contents.replace(/globalThis/mg, globalName);
    // contents = contents.replace('(() => {', `export function ${providerName} (${globalName}={crypto:globalThis.crypto,performance:globalThis.performance,TextEncoder:globalThis.TextEncoder,TextDecoder:globalThis.TextDecoder}) {`)
    // contents = contents.replace(`${globalName}.Go`, "const Go");
    // contents = contents.replace(`fs.`, `${globalName}.fs.`);
    // contents = contents.replace(`crypto.`, `${globalName}.crypto.`);
    // contents = contents.replace(`process.`, `${globalName}.process.`);
    // contents = contents.replace(`performance.`, `${globalName}.performance.`);
    // contents = contents.replace(`new TextEncoder.`, `new ${globalName}.TextEncoder`);
    // contents = contents.replace(`new TextDecoder.`, `new ${globalName}.TextDecoder`);

    const endOfFile = `
  globalThis.process.env = {...globalThis.process.env, AWS_PROFILE: "default"};
  return new Go();
}
globalThis.${providerName} = ${providerName};`;
    contents = contents.replace("})();", endOfFile);

//     contents += `
// ${providerName}.client = {};
// ${providerName}.ec2 = {};
// `;
    await fs.promises.writeFile(path.resolve(__dirname, "..", "..", outDir, "wasm_exec.js"), contents);
  }

  copyWasmExec(wasmExec, outDir).then(console.log("done"));
}

async function copyFiles(outDir) {
  const files = [path.join("misc", "index.html"), path.join("misc", "index.js")]
  for(const src of files) {
    await fs.copyFile(src, path.join(outDir, path.basename(src)))
  }
}

async function buildWeb(dir) {
  const outDir = path.join(dir, "web")
  await $`mkdir -p ${outDir}`
  await $`GOOS=js GOARCH=wasm go build -o ${outDir}/main.wasm -v`
  await patchWasmExec(outDir)
  await copyFiles(outDir)
}

async function buildV1(dir) {
  const outDir = path.join(dir, "sdkv1")
  await $`mkdir -p ${outDir}`
  await $`rm -rf ${outDir}/*`
  await $`cp -rL aws/aws-sdk-go/* ${outDir}`
}

async function build() {
  // compile
  const outBaseDir = "dist";
  await buildWeb(outBaseDir);
  await buildV1(outBaseDir);
}

build()
  .then(() => {
    console.log(chalk.bgGreen("build done"));
  })
  .catch((e) => {
    console.log(chalk.bgRed("build failed"));
    console.error(e);
    process.exit(1);
  });
