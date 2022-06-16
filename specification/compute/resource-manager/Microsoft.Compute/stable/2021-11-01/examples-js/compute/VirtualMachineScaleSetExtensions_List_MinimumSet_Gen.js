const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function virtualMachineScaleSetExtensionsListMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const vmScaleSetName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineScaleSetExtensions.list(
    resourceGroupName,
    vmScaleSetName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualMachineScaleSetExtensionsListMinimumSetGen().catch(console.error);
