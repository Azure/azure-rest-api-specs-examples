const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Endpoint resource, which represents a data transfer source or destination.
 *
 * @summary Creates or updates an Endpoint resource, which represents a data transfer source or destination.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/Endpoints_CreateOrUpdate.json
 */
async function endpointsCreateOrUpdate() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "11111111-2222-3333-4444-555555555555";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const endpointName = "examples-endpointName";
  const endpoint = {
    properties: {
      description: "Example Storage Container Endpoint Description",
      blobContainerName: "examples-blobContainerName",
      endpointType: "AzureStorageBlobContainer",
      storageAccountResourceId:
        "/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examples-storageAccountName/",
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
