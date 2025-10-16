const { OracleDatabaseManagementClient } = require("@azure/arm-oracledatabase");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list Private IP Addresses by the provided filter
 *
 * @summary list Private IP Addresses by the provided filter
 * x-ms-original-file: 2025-09-01/CloudVmClusters_ListPrivateIpAddresses_MinimumSet_Gen.json
 */
async function listPrivateIPAddressesForVMClusterGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new OracleDatabaseManagementClient(credential, subscriptionId);
  const result = await client.cloudVmClusters.listPrivateIpAddresses(
    "rgopenapi",
    "cloudvmcluster1",
    { subnetId: "ocid1..aaaaaa", vnicId: "ocid1..aaaaa" },
  );
  console.log(result);
}
