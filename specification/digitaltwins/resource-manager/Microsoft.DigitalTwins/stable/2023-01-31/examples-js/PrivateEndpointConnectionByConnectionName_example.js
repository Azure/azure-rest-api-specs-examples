const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get private endpoint connection properties for the given private endpoint.
 *
 * @summary Get private endpoint connection properties for the given private endpoint.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/PrivateEndpointConnectionByConnectionName_example.json
 */
async function getPrivateEndpointConnectionPropertiesForTheGivenPrivateEndpoint() {
  const subscriptionId =
    process.env["DIGITALTWINS_SUBSCRIPTION_ID"] || "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = process.env["DIGITALTWINS_RESOURCE_GROUP"] || "resRg";
  const resourceName = "myDigitalTwinsService";
  const privateEndpointConnectionName = "myPrivateConnection";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.get(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName
  );
  console.log(result);
}
