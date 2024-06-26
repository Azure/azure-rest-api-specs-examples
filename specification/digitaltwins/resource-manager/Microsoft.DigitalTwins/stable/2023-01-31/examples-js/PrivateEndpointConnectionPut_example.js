const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the status of a private endpoint connection with the given name.
 *
 * @summary Update the status of a private endpoint connection with the given name.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/PrivateEndpointConnectionPut_example.json
 */
async function updateTheStatusOfAPrivateEndpointConnectionWithTheGivenName() {
  const subscriptionId =
    process.env["DIGITALTWINS_SUBSCRIPTION_ID"] || "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = process.env["DIGITALTWINS_RESOURCE_GROUP"] || "resRg";
  const resourceName = "myDigitalTwinsService";
  const privateEndpointConnectionName = "myPrivateConnection";
  const privateEndpointConnection = {
    properties: {
      privateLinkServiceConnectionState: {
        description: "Approved by johndoe@company.com.",
        status: "Approved",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName,
    privateEndpointConnection
  );
  console.log(result);
}
