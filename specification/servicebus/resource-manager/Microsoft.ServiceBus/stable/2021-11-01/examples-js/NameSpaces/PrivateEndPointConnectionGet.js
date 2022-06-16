const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a description for the specified Private Endpoint Connection.
 *
 * @summary Gets a description for the specified Private Endpoint Connection.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/PrivateEndPointConnectionGet.json
 */
async function nameSpacePrivateEndPointConnectionGet() {
  const subscriptionId = "subID";
  const resourceGroupName = "SDK-ServiceBus-4794";
  const namespaceName = "sdk-Namespace-5828";
  const privateEndpointConnectionName = "privateEndpointConnectionName";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.get(
    resourceGroupName,
    namespaceName,
    privateEndpointConnectionName
  );
  console.log(result);
}

nameSpacePrivateEndPointConnectionGet().catch(console.error);
