const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a dedicated host.
 *
 * @summary Delete a dedicated host.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/dedicatedHostExamples/DedicatedHosts_Delete_MinimumSet_Gen.json
 */
async function dedicatedHostsDeleteMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const hostGroupName = "aaaaaaaaaaaaaaa";
  const hostName = "aaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHosts.beginDeleteAndWait(
    resourceGroupName,
    hostGroupName,
    hostName
  );
  console.log(result);
}

dedicatedHostsDeleteMinimumSetGen().catch(console.error);
