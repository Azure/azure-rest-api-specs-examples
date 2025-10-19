const { StorageMoverClient } = require("@azure/arm-storagemover");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates properties for an Endpoint resource. Properties not specified in the request body will be unchanged.
 *
 * @summary updates properties for an Endpoint resource. Properties not specified in the request body will be unchanged.
 * x-ms-original-file: 2025-07-01/Endpoints_Update_NfsMount.json
 */
async function endpointsUpdateNfsMount() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "60bcfc77-6589-4da2-b7fd-f9ec9322cf95";
  const client = new StorageMoverClient(credential, subscriptionId);
  const result = await client.endpoints.update(
    "examples-rg",
    "examples-storageMoverName",
    "examples-endpointName",
    {
      properties: {
        description: "Updated Endpoint Description",
        endpointType: "NfsMount",
      },
    },
  );
  console.log(result);
}
