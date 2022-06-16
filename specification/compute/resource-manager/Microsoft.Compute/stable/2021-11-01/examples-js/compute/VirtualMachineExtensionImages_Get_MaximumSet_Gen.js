const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineExtensionImagesGetMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaa";
  const type = "aaaaaaaaaaaaaaaaaa";
  const version = "aaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.get(
    location,
    publisherName,
    type,
    version
  );
  console.log(result);
}

virtualMachineExtensionImagesGetMaximumSetGen().catch(console.error);
