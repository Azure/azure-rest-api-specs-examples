const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the available PrivateEndpointConnections within a namespace.
 *
 * @summary Gets the available PrivateEndpointConnections within a namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/PrivateEndPointConnectionList.json
 */
async function nameSpaceCreate() {
  const subscriptionId = process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "subID";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "SDK-ServiceBus-4794";
  const namespaceName = "sdk-Namespace-5828";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndpointConnections.list(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
