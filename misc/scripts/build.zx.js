
async function patchWasmExec() {
  const goroot = await $`go env GOROOT`;
  const wasmExec = path.join(goroot.stdout.replace("\n", ""), "misc", "wasm", "wasm_exec.js");
  if (!wasmExec.endsWith("wasm_exec.js")) {
    throw new Error(
      "You must pass the absolute path to wasm_exec as the last argument to this script.\nnode ./scripts/wasm-exec.js $(go env GOROOT)/misc/wasm/wasm_exec.js"
    );
  }

  async function copyWasmExec(wasmExec) {
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
    await fs.promises.writeFile(path.resolve(__dirname, "..", "..", "dist", "wasm_exec.js"), contents);
  }

  copyWasmExec(wasmExec).then(console.log("done"));
}

async function copyFiles() {
  const files = [path.join("misc", "index.html"), path.join("misc", "index.js")]
  for(const src of files) {
    await fs.copyFile(src, path.join("dist", path.basename(src)))
  }
}


async function build() {
  // compile
  await $`mkdir -p dist`
  await $`GOOS=js GOARCH=wasm go build -o main.wasm -v`
  await $`mv main.wasm dist`
  await patchWasmExec()
  await copyFiles()
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
