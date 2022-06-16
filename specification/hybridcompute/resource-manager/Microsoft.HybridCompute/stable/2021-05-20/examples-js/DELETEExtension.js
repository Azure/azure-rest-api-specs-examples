const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteAMachineExtension() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const machineName = "myMachine";
  const extensionName = "MMA";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.machineExtensions.beginDeleteAndWait(
    resourceGroupName,
    machineName,
    extensionName
  );
  console.log(result);
}

deleteAMachineExtension().catch(console.error);
