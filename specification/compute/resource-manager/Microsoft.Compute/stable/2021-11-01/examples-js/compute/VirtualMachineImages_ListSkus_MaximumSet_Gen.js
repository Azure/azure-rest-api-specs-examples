const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineImagesListSkusMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaa";
  const publisherName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const offer = "aaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.listSkus(location, publisherName, offer);
  console.log(result);
}

virtualMachineImagesListSkusMaximumSetGen().catch(console.error);
