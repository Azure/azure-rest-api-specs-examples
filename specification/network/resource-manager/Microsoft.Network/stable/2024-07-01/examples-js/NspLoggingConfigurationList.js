const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Lists the NSP logging configuration.
 *
 * @summary Lists the NSP logging configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NspLoggingConfigurationList.json
 */
async function nspLoggingConfigurationList() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkSecurityPerimeterName = "nsp1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.networkSecurityPerimeterLoggingConfigurations.list(
    resourceGroupName,
    networkSecurityPerimeterName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
