const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a DICOM Service resource with the specified parameters.
 *
 * @summary Creates or updates a DICOM Service resource with the specified parameters.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_Create.json
 */
async function createOrUpdateADicomService() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const dicomServiceName = "blue";
  const dicomservice = { location: "westus" };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.dicomServices.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    dicomServiceName,
    dicomservice
  );
  console.log(result);
}

createOrUpdateADicomService().catch(console.error);
