const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function listMachinesByResourceGroup() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.machines.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listMachinesByResourceGroup().catch(console.error);
