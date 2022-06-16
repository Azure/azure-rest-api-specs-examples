const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves information about a dedicated host group.
 *
 * @summary Retrieves information about a dedicated host group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroup_Get_UltraSSDEnabledDedicatedHostGroup.json
 */
async function createAnUltraSsdEnabledDedicatedHostGroup() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const hostGroupName = "myDedicatedHostGroup";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHostGroups.get(resourceGroupName, hostGroupName);
  console.log(result);
}

createAnUltraSsdEnabledDedicatedHostGroup().catch(console.error);
