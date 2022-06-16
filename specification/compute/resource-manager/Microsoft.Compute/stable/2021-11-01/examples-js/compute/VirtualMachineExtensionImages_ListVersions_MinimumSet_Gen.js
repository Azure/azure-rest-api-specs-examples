const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineExtensionImagesListVersionsMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const type = "aaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensionImages.listVersions(
    location,
    publisherName,
    type
  );
  console.log(result);
}

virtualMachineExtensionImagesListVersionsMinimumSetGen().catch(console.error);
