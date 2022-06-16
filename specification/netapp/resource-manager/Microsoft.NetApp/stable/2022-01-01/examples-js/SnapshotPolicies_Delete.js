const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete snapshot policy
 *
 * @summary Delete snapshot policy
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/SnapshotPolicies_Delete.json
 */
async function snapshotPoliciesDelete() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = "resourceGroup";
  const accountName = "accountName";
  const snapshotPolicyName = "snapshotPolicyName";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.snapshotPolicies.beginDeleteAndWait(
    resourceGroupName,
    accountName,
    snapshotPolicyName
  );
  console.log(result);
}

snapshotPoliciesDelete().catch(console.error);
