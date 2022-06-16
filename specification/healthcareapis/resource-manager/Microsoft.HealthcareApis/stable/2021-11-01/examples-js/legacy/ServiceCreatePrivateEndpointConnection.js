const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of the specified private endpoint connection associated with the service.
 *
 * @summary Update the state of the specified private endpoint connection associated with the service.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/legacy/ServiceCreatePrivateEndpointConnection.json
 */
async function privateEndpointConnectionCreateOrUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgname";
  const resourceName = "service1";
  const privateEndpointConnectionName = "myConnection";
  const properties = {
    privateLinkServiceConnectionState: {
      description: "Auto-Approved",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName,
    properties
  );
  console.log(result);
}

privateEndpointConnectionCreateOrUpdate().catch(console.error);
