const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a snapshot Policy
 *
 * @summary Get a snapshot Policy
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-07-01/examples/SnapshotPolicies_Get.json
 */
async function snapshotPoliciesGet() {
  const subscriptionId =
    process.env["NETAPP_SUBSCRIPTION_ID"] || "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const resourceGroupName = process.env["NETAPP_RESOURCE_GROUP"] || "myRG";
  const accountName = "account1";
  const snapshotPolicyName = "snapshotPolicyName";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.snapshotPolicies.get(
    resourceGroupName,
    accountName,
    snapshotPolicyName,
  );
  console.log(result);
}
