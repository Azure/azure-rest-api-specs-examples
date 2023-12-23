const { HybridNetworkManagementClient } = require("@azure/arm-hybridnetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available artifacts in the parent Artifact Store.
 *
 * @summary Lists all the available artifacts in the parent Artifact Store.
 * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/PureProxyArtifact/ArtifactList.json
 */
async function listArtifactsUnderAnArtifactStore() {
  const subscriptionId = process.env["HYBRIDNETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDNETWORK_RESOURCE_GROUP"] || "TestResourceGroup";
  const publisherName = "TestPublisher";
  const artifactStoreName = "TestArtifactStoreName";
  const credential = new DefaultAzureCredential();
  const client = new HybridNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.proxyArtifact.list(
    resourceGroupName,
    publisherName,
    artifactStoreName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
