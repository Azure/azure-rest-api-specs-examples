const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get volumes associated with snapshot policy
 *
 * @summary Get volumes associated with snapshot policy
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-03-01/examples/SnapshotPolicies_ListVolumes.json
 */
async function snapshotPoliciesListVolumes() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "myRG";
  const accountName = "account1";
  const snapshotPolicyName = "snapshotPolicyName";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.snapshotPolicies.listVolumes(
    resourceGroupName,
    accountName,
    snapshotPolicyName
  );
  console.log(result);
}

snapshotPoliciesListVolumes().catch(console.error);
