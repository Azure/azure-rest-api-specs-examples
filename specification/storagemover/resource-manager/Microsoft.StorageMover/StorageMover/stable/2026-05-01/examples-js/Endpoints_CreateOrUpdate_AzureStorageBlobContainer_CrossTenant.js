const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Endpoint resource, which represents a data transfer source or destination.
 *
 * @summary creates or updates an Endpoint resource, which represents a data transfer source or destination.
 * x-ms-original-file: 2026-05-01/Endpoints_CreateOrUpdate_AzureStorageBlobContainer_CrossTenant.json
 */
async function endpointsCreateOrUpdateAzureStorageBlobContainerCrossTenant() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.endpoints.createOrUpdate(
    "examples-rg",
    "examples-storageMoverName",
    "examples-endpointName",
    {
      properties: {
        description: "Cross-tenant source Storage Blob Container Endpoint",
        blobContainerName: "examples-blobcontainer",
        endpointType: "AzureStorageBlobContainer",
        endpointKind: "Source",
        storageAccountResourceId:
          "/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesourcesa",
        enableCrossTenantTransfer: true,
        allowedStorageAccounts: [
          "/subscriptions/0a2b3c4d-5e6f-7081-92a3-b4c5d6e7f809/resourceGroups/partner-rg/providers/Microsoft.Storage/storageAccounts/partnertargetsa",
        ],
      },
    },
  );
  console.log(result);
}
