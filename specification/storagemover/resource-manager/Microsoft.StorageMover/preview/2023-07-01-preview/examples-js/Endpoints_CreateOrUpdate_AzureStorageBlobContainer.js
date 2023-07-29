const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Endpoint resource, which represents a data transfer source or destination.
 *
 * @summary Creates or updates an Endpoint resource, which represents a data transfer source or destination.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Endpoints_CreateOrUpdate_AzureStorageBlobContainer.json
 */
async function endpointsCreateOrUpdateAzureStorageBlobContainer() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const endpointName = "examples-endpointName";
  const endpoint = {
    properties: {
      description: "Example Storage Blob Container Endpoint Description",
      blobContainerName: "examples-blobcontainer",
      endpointType: "AzureStorageBlobContainer",
      storageAccountResourceId:
        "/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examplesa",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.endpoints.createOrUpdate(
    resourceGroupName,
    storageMoverName,
    endpointName,
    endpoint
  );
  console.log(result);
}
