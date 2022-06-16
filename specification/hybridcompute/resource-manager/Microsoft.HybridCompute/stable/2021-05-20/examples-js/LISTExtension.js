const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getAllMachineExtensions() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const machineName = "myMachine";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.machineExtensions.list(resourceGroupName, machineName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllMachineExtensions().catch(console.error);
