const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetsListMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineScaleSets.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualMachineScaleSetsListMinimumSetGen().catch(console.error);
