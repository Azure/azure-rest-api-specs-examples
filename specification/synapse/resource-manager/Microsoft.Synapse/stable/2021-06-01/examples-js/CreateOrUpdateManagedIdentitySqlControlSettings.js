const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update Managed Identity Sql Control Settings
 *
 * @summary Create or update Managed Identity Sql Control Settings
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateManagedIdentitySqlControlSettings.json
 */
async function createOrUpdateManagedIdentitySqlControlSettings() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "resourceGroup1";
  const workspaceName = "workspace1";
  const managedIdentitySqlControlSettings = {
    grantSqlControlToManagedIdentity: { desiredState: "Enabled" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.workspaceManagedIdentitySqlControlSettings.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    managedIdentitySqlControlSettings
  );
  console.log(result);
}
