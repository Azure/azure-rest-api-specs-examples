const { QumuloStorage } = require("@azure/arm-qumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete a FileSystemResource
 *
 * @summary delete a FileSystemResource
 * x-ms-original-file: 2026-04-16/FileSystems_Delete_MaximumSet_Gen.json
 */
async function fileSystemsDeleteMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "C9CC2D2A-5AA0-4839-A85F-18491F2D244A";
  const client = new QumuloStorage(credential, subscriptionId);
  await client.fileSystems.delete("rgQumulo", "qumulo-fs-01");
}
