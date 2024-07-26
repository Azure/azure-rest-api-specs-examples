const { QumuloStorage } = require("@azure/arm-qumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a FileSystemResource
 *
 * @summary Delete a FileSystemResource
 * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/FileSystems_Delete_MaximumSet_Gen.json
 */
async function fileSystemsDelete() {
  const subscriptionId =
    process.env["LIFTRQUMULO_SUBSCRIPTION_ID"] || "382E8C7A-AC80-4D70-8580-EFE99537B9B7";
  const resourceGroupName = process.env["LIFTRQUMULO_RESOURCE_GROUP"] || "rgQumulo";
  const fileSystemName = "xoschzkccroahrykedlvbbnsddq";
  const credential = new DefaultAzureCredential();
  const client = new QumuloStorage(credential, subscriptionId);
  const result = await client.fileSystems.beginDeleteAndWait(resourceGroupName, fileSystemName);
  console.log(result);
}
