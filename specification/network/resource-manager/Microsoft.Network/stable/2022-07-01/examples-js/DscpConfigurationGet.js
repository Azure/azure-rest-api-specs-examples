const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a DSCP Configuration.
 *
 * @summary Gets a DSCP Configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/DscpConfigurationGet.json
 */
async function getDscpConfiguration() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const dscpConfigurationName = "mydscpConfig";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.dscpConfigurationOperations.get(
    resourceGroupName,
    dscpConfigurationName
  );
  console.log(result);
}

getDscpConfiguration().catch(console.error);
