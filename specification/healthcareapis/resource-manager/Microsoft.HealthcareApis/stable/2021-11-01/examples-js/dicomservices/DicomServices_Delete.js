const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a DICOM Service.
 *
 * @summary Deletes a DICOM Service.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_Delete.json
 */
async function deleteADicomservice() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const dicomServiceName = "blue";
  const workspaceName = "workspace1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.dicomServices.beginDeleteAndWait(
    resourceGroupName,
    dicomServiceName,
    workspaceName
  );
  console.log(result);
}

deleteADicomservice().catch(console.error);
