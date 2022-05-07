Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-healthcareapis_2.1.0/sdk/healthcareapis/arm-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the properties of the specified DICOM Service.
 *
 * @summary Gets the properties of the specified DICOM Service.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/dicomservices/DicomServices_Get.json
 */
async function getADicomservice() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const workspaceName = "workspace1";
  const dicomServiceName = "blue";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.dicomServices.get(resourceGroupName, workspaceName, dicomServiceName);
  console.log(result);
}

getADicomservice().catch(console.error);
```
