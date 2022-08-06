const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Transform.
 *
 * @summary Deletes a Transform.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-delete.json
 */
async function deleteATransform() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contosoresources";
  const accountName = "contosomedia";
  const transformName = "sampleTransform";
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.transforms.delete(resourceGroupName, accountName, transformName);
  console.log(result);
}

deleteATransform().catch(console.error);
