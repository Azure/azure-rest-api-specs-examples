const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all PrivateEndpointConnections in the PrivateLinkHub
 *
 * @summary Get all PrivateEndpointConnections in the PrivateLinkHub
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/PrivateEndpointConnectionsPrivateLinkHub_List.json
 */
async function getAPrivateLinkHub() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "48b08652-d7a1-4d52-b13f-5a2471dce57b";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "gh-res-grp";
  const privateLinkHubName = "pe0";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnectionsPrivateLinkHub.list(
    resourceGroupName,
    privateLinkHubName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
