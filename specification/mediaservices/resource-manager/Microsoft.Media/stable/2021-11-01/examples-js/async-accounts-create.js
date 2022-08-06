const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Media Services account
 *
 * @summary Creates or updates a Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/async-accounts-create.json
 */
async function createAMediaServicesAccount() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contososports";
  const parameters = {
    encryption: {
      type: "CustomerKey",
      identity: {
        useSystemAssignedIdentity: false,
        userAssignedIdentity:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
      },
    },
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourceGroups/contoso/providers/MicrosoftManagedIdentity/userAssignedIdentities/id1":
          {},
        "/subscriptions/00000000000000000000000000000000/resourceGroups/contoso/providers/MicrosoftManagedIdentity/userAssignedIdentities/id2":
          {},
      },
    },
    keyDelivery: { accessControl: { defaultAction: "Allow" } },
    location: "South Central US",
    publicNetworkAccess: "Enabled",
    storageAccounts: [
      {
        type: "Primary",
        id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.Storage/storageAccounts/contososportsstore",
        identity: {
          useSystemAssignedIdentity: false,
          userAssignedIdentity:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/contoso/providers/Microsoft.ManagedIdentity/userAssignedIdentities/id1",
        },
      },
    ],
    storageAuthentication: "ManagedIdentity",
    tags: { key1: "value1", key2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.mediaservices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

createAMediaServicesAccount().catch(console.error);
