const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Post a Network Manager Commit.
 *
 * @summary Post a Network Manager Commit.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkManagerCommitPost.json
 */
async function networkManageCommitPost() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "resoureGroupSample";
  const networkManagerName = "testNetworkManager";
  const parameters = {
    commitType: "SecurityAdmin",
    configurationIds: [
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resoureGroupSample/providers/Microsoft.Network/networkManagers/testNetworkManager/securityAdminConfigurations/SampleSecurityAdminConfig",
    ],
    targetLocations: ["useast"],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.networkManagerCommits.beginPostAndWait(
    resourceGroupName,
    networkManagerName,
    parameters
  );
  console.log(result);
}

networkManageCommitPost().catch(console.error);
