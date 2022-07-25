const { DataBoxManagementClient } = require("@azure/arm-databox");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the properties of an existing job.
 *
 * @summary Updates the properties of an existing job.
 * x-ms-original-file: specification/databox/resource-manager/Microsoft.DataBox/preview/2021-08-01-preview/examples/JobsPatchSystemAssignedToUserAssigned.json
 */
async function jobsPatchSystemAssignedToUserAssigned() {
  const subscriptionId = "fa68082f-8ff7-4a25-95c7-ce9da541242f";
  const resourceGroupName = "SdkRg9765";
  const jobName = "SdkJob2965";
  const jobResourceUpdateParameter = {
    identity: {
      type: "SystemAssigned,UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/fa68082f8ff74a2595c7Ce9da541242f/resourceGroups/akvenkat/providers/MicrosoftManagedIdentity/userAssignedIdentities/sdkIdentity":
          {},
      },
    },
    details: {
      keyEncryptionKey: {
        identityProperties: {
          type: "UserAssigned",
          userAssigned: {
            resourceId:
              "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.ManagedIdentity/userAssignedIdentities/sdkIdentity",
          },
        },
        kekType: "CustomerManaged",
        kekUrl: "https://sdkkeyvault.vault.azure.net/keys/SSDKEY/",
        kekVaultResourceID:
          "/subscriptions/fa68082f-8ff7-4a25-95c7-ce9da541242f/resourceGroups/akvenkat/providers/Microsoft.KeyVault/vaults/SDKKeyVault",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxManagementClient(credential, subscriptionId);
  const result = await client.jobs.beginUpdateAndWait(
    resourceGroupName,
    jobName,
    jobResourceUpdateParameter
  );
  console.log(result);
}

jobsPatchSystemAssignedToUserAssigned().catch(console.error);
