const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all FHIR Services for the given workspace
 *
 * @summary Lists all FHIR Services for the given workspace
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/fhirservices/FhirServices_List.json
 */
async function listFhirservices() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fhirServices.listByWorkspace(resourceGroupName, workspaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listFhirservices().catch(console.error);
