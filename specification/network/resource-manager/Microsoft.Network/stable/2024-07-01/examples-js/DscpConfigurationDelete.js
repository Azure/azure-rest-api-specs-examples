const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes a DSCP Configuration.
 *
 * @summary Deletes a DSCP Configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/DscpConfigurationDelete.json
 */
async function deleteDscpConfiguration() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const dscpConfigurationName = "mydscpConfig";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.dscpConfigurationOperations.beginDeleteAndWait(
    resourceGroupName,
    dscpConfigurationName,
  );
  console.log(result);
}
