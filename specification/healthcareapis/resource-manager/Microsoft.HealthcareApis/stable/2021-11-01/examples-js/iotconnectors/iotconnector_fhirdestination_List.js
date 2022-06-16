const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all FHIR destinations for the given IoT Connector
 *
 * @summary Lists all FHIR destinations for the given IoT Connector
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_List.json
 */
async function listIoTConnectors() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const iotConnectorName = "blue";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fhirDestinations.listByIotConnector(
    resourceGroupName,
    workspaceName,
    iotConnectorName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listIoTConnectors().catch(console.error);
