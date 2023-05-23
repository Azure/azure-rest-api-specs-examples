const { QumuloStorage } = require("@azure/arm-liftrqumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List FileSystemResource resources by resource group
 *
 * @summary List FileSystemResource resources by resource group
 * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_ListByResourceGroup_MinimumSet_Gen.json
 */
async function fileSystemsListByResourceGroupMinimumSetGen() {
  const subscriptionId = process.env["LIFTRQUMULO_SUBSCRIPTION_ID"] || "aaaaaaa";
  const resourceGroupName = process.env["LIFTRQUMULO_RESOURCE_GROUP"] || "rgQumulo";
  const credential = new DefaultAzureCredential();
  const client = new QumuloStorage(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fileSystems.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
