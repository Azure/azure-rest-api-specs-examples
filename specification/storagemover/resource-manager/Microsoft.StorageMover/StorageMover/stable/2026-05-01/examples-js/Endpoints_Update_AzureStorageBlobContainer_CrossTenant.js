const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates properties for an Endpoint resource. Properties not specified in the request body will be unchanged.
 *
 * @summary updates properties for an Endpoint resource. Properties not specified in the request body will be unchanged.
 * x-ms-original-file: 2026-05-01/Endpoints_Update_AzureStorageBlobContainer_CrossTenant.json
 */
async function endpointsUpdateAzureStorageBlobContainerCrossTenant() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.endpoints.update(
    "examples-rg",
    "examples-storageMoverName",
    "examples-endpointName",
    {
      properties: {
        description:
          "Updated cross-tenant source endpoint with refreshed allowed partner storage accounts",
        endpointType: "AzureStorageBlobContainer",
        enableCrossTenantTransfer: true,
        allowedStorageAccounts: [
          "/subscriptions/0a2b3c4d-5e6f-7081-92a3-b4c5d6e7f809/resourceGroups/partner-rg/providers/Microsoft.Storage/storageAccounts/partnertargetsa",
          "/subscriptions/0a2b3c4d-5e6f-7081-92a3-b4c5d6e7f809/resourceGroups/partner-rg/providers/Microsoft.Storage/storageAccounts/partnertargetsa2",
        ],
      },
    },
  );
  console.log(result);
}
