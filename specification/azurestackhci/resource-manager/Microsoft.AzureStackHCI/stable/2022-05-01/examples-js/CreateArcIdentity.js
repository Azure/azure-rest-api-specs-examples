const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Aad identity for arc settings.
 *
 * @summary Create Aad identity for arc settings.
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/CreateArcIdentity.json
 */
async function createArcIdentity() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "test-rg";
  const clusterName = "myCluster";
  const arcSettingName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.arcSettings.beginCreateIdentityAndWait(
    resourceGroupName,
    clusterName,
    arcSettingName
  );
  console.log(result);
}

createArcIdentity().catch(console.error);
