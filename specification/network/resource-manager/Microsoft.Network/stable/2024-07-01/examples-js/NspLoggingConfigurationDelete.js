const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes an NSP Logging configuration.
 *
 * @summary Deletes an NSP Logging configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NspLoggingConfigurationDelete.json
 */
async function nspLoggingConfigurationDelete() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkSecurityPerimeterName = "nsp1";
  const loggingConfigurationName = "instance";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkSecurityPerimeterLoggingConfigurations.delete(
    resourceGroupName,
    networkSecurityPerimeterName,
    loggingConfigurationName,
  );
  console.log(result);
}
