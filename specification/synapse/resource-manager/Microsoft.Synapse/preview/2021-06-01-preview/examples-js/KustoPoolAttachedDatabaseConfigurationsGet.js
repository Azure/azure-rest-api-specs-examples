const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns an attached database configuration.
 *
 * @summary Returns an attached database configuration.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsGet.json
 */
async function kustoPoolAttachedDatabaseConfigurationsGet() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789098";
  const workspaceName = "kustorptest";
  const kustoPoolName = "kustoclusterrptest4";
  const attachedDatabaseConfigurationName = "attachedDatabaseConfigurations1";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "kustorptest";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolAttachedDatabaseConfigurations.get(
    workspaceName,
    kustoPoolName,
    attachedDatabaseConfigurationName,
    resourceGroupName
  );
  console.log(result);
}
