const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteAMachine() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const machineName = "myMachine";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.machines.delete(resourceGroupName, machineName);
  console.log(result);
}

deleteAMachine().catch(console.error);
