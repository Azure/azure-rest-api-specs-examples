const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all service endpoint policy definitions in a service end point policy.
 *
 * @summary Gets all service endpoint policy definitions in a service end point policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/ServiceEndpointPolicyDefinitionList.json
 */
async function listServiceEndpointDefinitionsInServiceEndPointPolicy() {
  const subscriptionId = process.env["NETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const serviceEndpointPolicyName = "testPolicy";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serviceEndpointPolicyDefinitions.listByResourceGroup(
    resourceGroupName,
    serviceEndpointPolicyName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
