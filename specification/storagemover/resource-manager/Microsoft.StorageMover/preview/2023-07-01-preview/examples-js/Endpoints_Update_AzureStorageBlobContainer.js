const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates properties for an Endpoint resource. Properties not specified in the request body will be unchanged.
 *
 * @summary Updates properties for an Endpoint resource. Properties not specified in the request body will be unchanged.
 * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Endpoints_Update_AzureStorageBlobContainer.json
 */
async function endpointsUpdateAzureStorageBlobContainer() {
  const subscriptionId =
    process.env["STORAGEMOVER_SUBSCRIPTION_ID"] || "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const resourceGroupName = process.env["STORAGEMOVER_RESOURCE_GROUP"] || "examples-rg";
  const storageMoverName = "examples-storageMoverName";
  const endpointName = "examples-endpointName";
  const endpoint = {
    properties: {
      description: "Updated Endpoint Description",
      endpointType: "AzureStorageBlobContainer",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.endpoints.update(
    resourceGroupName,
    storageMoverName,
    endpointName,
    endpoint
  );
  console.log(result);
}
