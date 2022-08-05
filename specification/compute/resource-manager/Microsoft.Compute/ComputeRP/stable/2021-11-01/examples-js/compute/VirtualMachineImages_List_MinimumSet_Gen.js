const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineImagesListMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaa";
  const publisherName = "aaaaaaaaaaa";
  const offer = "aaaaaaaaaa";
  const skus = "aaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.list(location, publisherName, offer, skus);
  console.log(result);
}

virtualMachineImagesListMinimumSetGen().catch(console.error);
