const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a segment by id in a private cloud workload network.
 *
 * @summary Get a segment by id in a private cloud workload network.
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/WorkloadNetworks_GetSegments.json
 */
async function workloadNetworksGetSegment() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const segmentId = "segment1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.getSegment(
    resourceGroupName,
    privateCloudName,
    segmentId
  );
  console.log(result);
}

workloadNetworksGetSegment().catch(console.error);
