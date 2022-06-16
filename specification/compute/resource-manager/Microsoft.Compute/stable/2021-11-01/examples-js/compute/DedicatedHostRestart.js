const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function restartDedicatedHost() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const hostGroupName = "myDedicatedHostGroup";
  const hostName = "myHost";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHosts.beginRestartAndWait(
    resourceGroupName,
    hostGroupName,
    hostName
  );
  console.log(result);
}

restartDedicatedHost().catch(console.error);
