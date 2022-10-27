const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a placement policy in a private cloud cluster
 *
 * @summary Delete a placement policy in a private cloud cluster
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/PlacementPolicies_Delete.json
 */
async function placementPoliciesDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const clusterName = "cluster1";
  const placementPolicyName = "policy1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.placementPolicies.beginDeleteAndWait(
    resourceGroupName,
    privateCloudName,
    clusterName,
    placementPolicyName
  );
  console.log(result);
}

placementPoliciesDelete().catch(console.error);
