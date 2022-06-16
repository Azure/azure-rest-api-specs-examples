const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function listAllVirtualMachineImagesInASubscription() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.images.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllVirtualMachineImagesInASubscription().catch(console.error);
