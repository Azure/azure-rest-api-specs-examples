const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a placement policy by name in a private cloud cluster
 *
 * @summary Get a placement policy by name in a private cloud cluster
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/PlacementPolicies_Get.json
 */
async function placementPoliciesGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const clusterName = "cluster1";
  const placementPolicyName = "policy1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.placementPolicies.get(
    resourceGroupName,
    privateCloudName,
    clusterName,
    placementPolicyName
  );
  console.log(result);
}

placementPoliciesGet().catch(console.error);
