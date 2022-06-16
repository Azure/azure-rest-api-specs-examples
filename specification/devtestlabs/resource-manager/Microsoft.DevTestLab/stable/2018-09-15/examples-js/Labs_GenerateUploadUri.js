const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate a URI for uploading custom disk images to a Lab.
 *
 * @summary Generate a URI for uploading custom disk images to a Lab.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_GenerateUploadUri.json
 */
async function labsGenerateUploadUri() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const name = "{labName}";
  const generateUploadUriParameter = {
    blobName: "{blob-name}",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.labs.generateUploadUri(
    resourceGroupName,
    name,
    generateUploadUriParameter
  );
  console.log(result);
}

labsGenerateUploadUri().catch(console.error);
