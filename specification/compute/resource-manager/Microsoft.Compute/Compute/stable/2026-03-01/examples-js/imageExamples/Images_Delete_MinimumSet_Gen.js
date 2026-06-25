const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes an Image.
 *
 * @summary deletes an Image.
 * x-ms-original-file: 2026-03-01/imageExamples/Images_Delete_MinimumSet_Gen.json
 */
async function imageDeleteMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  await client.images.delete("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaaa");
}
