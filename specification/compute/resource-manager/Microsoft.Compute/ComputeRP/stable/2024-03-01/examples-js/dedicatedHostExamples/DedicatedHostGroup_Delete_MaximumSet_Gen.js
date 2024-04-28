const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a dedicated host group.
 *
 * @summary Delete a dedicated host group.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/dedicatedHostExamples/DedicatedHostGroup_Delete_MaximumSet_Gen.json
 */
async function dedicatedHostGroupDeleteMaximumSetGen() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "rgcompute";
  const hostGroupName = "a";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHostGroups.delete(resourceGroupName, hostGroupName);
  console.log(result);
}
