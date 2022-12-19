const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List private endpoint connection properties.
 *
 * @summary List private endpoint connection properties.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/PrivateEndpointConnectionsList_example.json
 */
async function listPrivateEndpointConnectionProperties() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.list(resourceGroupName, resourceName);
  console.log(result);
}

listPrivateEndpointConnectionProperties().catch(console.error);
