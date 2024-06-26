const ServiceFabricManagementClient = require("@azure-rest/arm-servicefabric").default;
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to If a target is not provided, it will get the minimum and maximum versions available from the current cluster version. If a target is given, it will provide the required path to get from the current cluster version to the target version.
 *
 * @summary If a target is not provided, it will get the minimum and maximum versions available from the current cluster version. If a target is given, it will provide the required path to get from the current cluster version to the target version.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ListUpgradableVersionsMinMax_example.json
 */
async function getMinimumAndMaximumCodeVersions() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const credential = new DefaultAzureCredential();
  const client = ServiceFabricManagementClient(credential);
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/listUpgradableVersions",
      subscriptionId,
      resourceGroupName,
      clusterName
    )
    .post();
  console.log(result);
}

getMinimumAndMaximumCodeVersions().catch(console.error);
