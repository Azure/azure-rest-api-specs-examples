const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Post to List of Network Manager Deployment Status.
 *
 * @summary Post to List of Network Manager Deployment Status.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerDeploymentStatusList.json
 */
async function networkManagerDeploymentStatusList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resoureGroupSample";
  const networkManagerName = "testNetworkManager";
  const parameters = {
    deploymentTypes: ["Connectivity", "AdminPolicy"],
    regions: ["eastus", "westus"],
    skipToken: "FakeSkipTokenCode",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkManagerDeploymentStatusOperations.list(
    resourceGroupName,
    networkManagerName,
    parameters
  );
  console.log(result);
}

networkManagerDeploymentStatusList().catch(console.error);
