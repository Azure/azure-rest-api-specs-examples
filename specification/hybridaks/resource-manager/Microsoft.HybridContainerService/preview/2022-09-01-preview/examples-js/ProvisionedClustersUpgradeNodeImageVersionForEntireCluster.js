const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Upgrading the node image version of a cluster applies the newest OS and runtime updates to the nodes.
 *
 * @summary Upgrading the node image version of a cluster applies the newest OS and runtime updates to the nodes.
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/ProvisionedClustersUpgradeNodeImageVersionForEntireCluster.json
 */
async function upgradeClusterNodeImageVersion() {
  const subscriptionId =
    process.env["HYBRIDCONTAINERSERVICE_SUBSCRIPTION_ID"] || "a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b";
  const resourceGroupName =
    process.env["HYBRIDCONTAINERSERVICE_RESOURCE_GROUP"] || "test-arcappliance-resgrp";
  const resourceName = "test-hybridakscluster";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential, subscriptionId);
  const result =
    await client.provisionedClustersOperations.beginUpgradeNodeImageVersionForEntireClusterAndWait(
      resourceGroupName,
      resourceName
    );
  console.log(result);
}
