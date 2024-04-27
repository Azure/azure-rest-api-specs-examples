const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the specified Iot Connector FHIR destination.
 *
 * @summary Gets the properties of the specified Iot Connector FHIR destination.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/iotconnectors/iotconnector_fhirdestination_Get.json
 */
async function getAnIoTConnectorDestination() {
  const subscriptionId = process.env["HEALTHCAREAPIS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HEALTHCAREAPIS_RESOURCE_GROUP"] || "testRG";
  const workspaceName = "workspace1";
  const iotConnectorName = "blue";
  const fhirDestinationName = "dest1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.iotConnectorFhirDestination.get(
    resourceGroupName,
    workspaceName,
    iotConnectorName,
    fhirDestinationName,
  );
  console.log(result);
}
