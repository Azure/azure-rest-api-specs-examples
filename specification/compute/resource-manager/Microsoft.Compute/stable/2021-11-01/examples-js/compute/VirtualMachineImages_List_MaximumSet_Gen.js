const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineImagesListMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaaaaa";
  const publisherName = "aaaaaa";
  const offer = "aaaaaaaaaaaaaaaa";
  const skus = "aaaaaaaaaaaaaaaaaaaaaaa";
  const expand = "aaaaaaaaaaaaaaaaaaaaaaaa";
  const top = 18;
  const orderby = "aa";
  const options = { expand: expand, top: top, orderby: orderby };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImages.list(
    location,
    publisherName,
    offer,
    skus,
    options
  );
  console.log(result);
}

virtualMachineImagesListMaximumSetGen().catch(console.error);
