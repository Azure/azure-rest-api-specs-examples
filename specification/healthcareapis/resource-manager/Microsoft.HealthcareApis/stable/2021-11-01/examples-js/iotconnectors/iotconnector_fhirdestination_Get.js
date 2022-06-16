const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the specified Iot Connector FHIR destination.
 *
 * @summary Gets the properties of the specified Iot Connector FHIR destination.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_Get.json
 */
async function getAnIoTConnectorDestination() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const iotConnectorName = "blue";
  const fhirDestinationName = "dest1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.iotConnectorFhirDestination.get(
    resourceGroupName,
    workspaceName,
    iotConnectorName,
    fhirDestinationName
  );
  console.log(result);
}

getAnIoTConnectorDestination().catch(console.error);
