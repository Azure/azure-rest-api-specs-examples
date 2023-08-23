const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Maps Account. Only a subset of the parameters may be updated after creation, such as Sku, Tags, Properties.
 *
 * @summary Updates a Maps Account. Only a subset of the parameters may be updated after creation, such as Sku, Tags, Properties.
 * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/UpdateAccountEncryption.json
 */
async function updateAccountEncryption() {
  const subscriptionId =
    process.env["MAPS_SUBSCRIPTION_ID"] || "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const resourceGroupName = process.env["MAPS_RESOURCE_GROUP"] || "myResourceGroup";
  const accountName = "myMapsAccount";
  const mapsAccountUpdateParameters = {
    encryption: {
      customerManagedKeyEncryption: {
        keyEncryptionKeyIdentity: {
          identityType: "systemAssignedIdentity",
          userAssignedIdentityResourceId: undefined,
        },
        keyEncryptionKeyUrl: "https://contosovault.vault.azure.net/keys/contosokek",
      },
    },
    identity: {
      type: "SystemAssigned",
      userAssignedIdentities: {
        "/subscriptions/21a9967aE8a94656A70b96ff1c4d05a0/resourceGroups/myResourceGroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/identityName":
          {},
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.accounts.update(
    resourceGroupName,
    accountName,
    mapsAccountUpdateParameters
  );
  console.log(result);
}
