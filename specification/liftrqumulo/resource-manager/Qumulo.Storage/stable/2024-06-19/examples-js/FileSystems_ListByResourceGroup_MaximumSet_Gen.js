const { QumuloStorage } = require("@azure/arm-qumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List FileSystemResource resources by resource group
 *
 * @summary List FileSystemResource resources by resource group
 * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2024-06-19/examples/FileSystems_ListByResourceGroup_MaximumSet_Gen.json
 */
async function fileSystemsListByResourceGroup() {
  const subscriptionId =
    process.env["LIFTRQUMULO_SUBSCRIPTION_ID"] || "382E8C7A-AC80-4D70-8580-EFE99537B9B7";
  const resourceGroupName = process.env["LIFTRQUMULO_RESOURCE_GROUP"] || "rgQumulo";
  const credential = new DefaultAzureCredential();
  const client = new QumuloStorage(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fileSystems.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
