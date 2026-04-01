const { AzureMapsManagementClient } = require("@azure/arm-maps");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Maps Account. A Maps Account holds the keys which allow access to the Maps REST APIs.
 *
 * @summary create or update a Maps Account. A Maps Account holds the keys which allow access to the Maps REST APIs.
 * x-ms-original-file: 2025-10-01-preview/CreateAccountManagedIdentity.json
 */
async function createAccountWithManagedIdentities() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "21a9967a-e8a9-4656-a70b-96ff1c4d05a0";
  const client = new AzureMapsManagementClient(credential, subscriptionId);
  const result = await client.accounts.createOrUpdate("myResourceGroup", "myMapsAccount", {
    identity: {
      type: "SystemAssigned, UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identityName":
          {},
      },
    },
    kind: "Gen2",
    location: "eastus",
    properties: {
      disableLocalAuth: false,
      linkedResources: [
        {
          id: "/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Storage/accounts/mystorageacc",
          uniqueName: "myBatchStorageAccount",
        },
        {
          id: "/subscriptions/21a9967a-e8a9-4656-a70b-96ff1c4d05a0/resourceGroups/myResourceGroup/providers/Microsoft.Storage/accounts/mystorageacc",
          uniqueName: "myBlobDataSource",
        },
      ],
    },
    sku: { name: "G2" },
    tags: { test: "true" },
  });
  console.log(result);
}
