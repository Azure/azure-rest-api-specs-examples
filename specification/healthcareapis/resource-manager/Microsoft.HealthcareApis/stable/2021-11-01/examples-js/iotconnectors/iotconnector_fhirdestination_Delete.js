const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an IoT Connector FHIR destination.
 *
 * @summary Deletes an IoT Connector FHIR destination.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_Delete.json
 */
async function deleteAnIoTConnectorDestination() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const iotConnectorName = "blue";
  const fhirDestinationName = "dest1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.iotConnectorFhirDestination.beginDeleteAndWait(
    resourceGroupName,
    workspaceName,
    iotConnectorName,
    fhirDestinationName
  );
  console.log(result);
}

deleteAnIoTConnectorDestination().catch(console.error);
