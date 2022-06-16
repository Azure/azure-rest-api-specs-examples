const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the available PrivateEndpointConnections within a namespace.
 *
 * @summary Gets the available PrivateEndpointConnections within a namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/PrivateEndPointConnectionList.json
 */
async function nameSpaceCreate() {
  const subscriptionId = "subID";
  const resourceGroupName = "SDK-EventHub-4794";
  const namespaceName = "sdk-Namespace-5828";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

nameSpaceCreate().catch(console.error);
