const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a DSCP Configuration.
 *
 * @summary Deletes a DSCP Configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/DscpConfigurationDelete.json
 */
async function deleteDscpConfiguration() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const dscpConfigurationName = "mydscpConfig";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.dscpConfigurationOperations.beginDeleteAndWait(
    resourceGroupName,
    dscpConfigurationName
  );
  console.log(result);
}

deleteDscpConfiguration().catch(console.error);
