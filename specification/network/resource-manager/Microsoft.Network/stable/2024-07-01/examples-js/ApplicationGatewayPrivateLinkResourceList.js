const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists all private link resources on an application gateway.
 *
 * @summary Lists all private link resources on an application gateway.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ApplicationGatewayPrivateLinkResourceList.json
 */
async function listsAllPrivateLinkResourcesOnApplicationGateway() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const applicationGatewayName = "appgw";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.applicationGatewayPrivateLinkResources.list(
    resourceGroupName,
    applicationGatewayName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
