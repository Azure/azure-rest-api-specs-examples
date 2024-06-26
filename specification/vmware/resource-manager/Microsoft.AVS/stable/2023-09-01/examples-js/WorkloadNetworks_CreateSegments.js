const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a WorkloadNetworkSegment
 *
 * @summary Create a WorkloadNetworkSegment
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-09-01/examples/WorkloadNetworks_CreateSegments.json
 */
async function workloadNetworksCreateSegments() {
  const subscriptionId =
    process.env["AVS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["AVS_RESOURCE_GROUP"] || "group1";
  const privateCloudName = "cloud1";
  const segmentId = "segment1";
  const workloadNetworkSegment = {
    connectedGateway: "/infra/tier-1s/gateway",
    displayName: "segment1",
    revision: 1,
    subnet: {
      dhcpRanges: ["40.20.0.0-40.20.0.1"],
      gatewayAddress: "40.20.20.20/16",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.workloadNetworks.beginCreateSegmentsAndWait(
    resourceGroupName,
    privateCloudName,
    segmentId,
    workloadNetworkSegment,
  );
  console.log(result);
}
