const { AzureTrafficCollectorClient } = require("@azure/arm-networkfunction");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return list of Azure Traffic Collectors in a Resource Group
 *
 * @summary Return list of Azure Traffic Collectors in a Resource Group
 * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/AzureTrafficCollectorsByResourceGroupList.json
 */
async function listOfTrafficCollectorsByResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new AzureTrafficCollectorClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.azureTrafficCollectorsByResourceGroup.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listOfTrafficCollectorsByResourceGroup().catch(console.error);
