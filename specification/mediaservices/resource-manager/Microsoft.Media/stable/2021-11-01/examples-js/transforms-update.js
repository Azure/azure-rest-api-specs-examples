const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Transform.
 *
 * @summary Updates a Transform.
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-update.json
 */
async function updateATransform() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contosoresources";
  const accountName = "contosomedia";
  const transformName = "transformToUpdate";
  const parameters = {
    description: "Example transform to illustrate update.",
    outputs: [
      {
        preset: {
          odataType: "#Microsoft.Media.BuiltInStandardEncoderPreset",
          presetName: "H264MultipleBitrate720p",
        },
        relativePriority: "High",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.transforms.update(
    resourceGroupName,
    accountName,
    transformName,
    parameters
  );
  console.log(result);
}

updateATransform().catch(console.error);
