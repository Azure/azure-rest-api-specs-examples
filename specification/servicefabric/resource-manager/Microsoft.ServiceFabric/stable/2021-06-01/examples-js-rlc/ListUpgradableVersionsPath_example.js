const ServiceFabricManagementClient = require("@azure-rest/arm-servicefabric").default;
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to If a target is not provided, it will get the minimum and maximum versions available from the current cluster version. If a target is given, it will provide the required path to get from the current cluster version to the target version.
 *
 * @summary If a target is not provided, it will get the minimum and maximum versions available from the current cluster version. If a target is given, it will provide the required path to get from the current cluster version to the target version.
 * x-ms-original-file: specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ListUpgradableVersionsPath_example.json
 */
async function getUpgradePath() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resRg";
  const clusterName = "myCluster";
  const options = {
    body: {
      targetVersion: "7.2.432.9590",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = ServiceFabricManagementClient(credential);
  const result = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceFabric/clusters/{clusterName}/listUpgradableVersions",
      subscriptionId,
      resourceGroupName,
      clusterName
    )
    .post(options);
  console.log(result);
}

getUpgradePath().catch(console.error);
