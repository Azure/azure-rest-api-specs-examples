const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all application gateways in a resource group.
 *
 * @summary Lists all application gateways in a resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ApplicationGatewayList.json
 */
async function listsAllApplicationGatewaysInAResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationGateways.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsAllApplicationGatewaysInAResourceGroup().catch(console.error);
