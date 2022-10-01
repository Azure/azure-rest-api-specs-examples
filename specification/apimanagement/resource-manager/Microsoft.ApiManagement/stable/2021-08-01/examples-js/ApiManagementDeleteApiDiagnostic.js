const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified Diagnostic from an API.
 *
 * @summary Deletes the specified Diagnostic from an API.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteApiDiagnostic.json
 */
async function apiManagementDeleteApiDiagnostic() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceName = "apimService1";
  const apiId = "57d1f7558aa04f15146d9d8a";
  const diagnosticId = "applicationinsights";
  const ifMatch = "*";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiDiagnostic.delete(
    resourceGroupName,
    serviceName,
    apiId,
    diagnosticId,
    ifMatch
  );
  console.log(result);
}

apiManagementDeleteApiDiagnostic().catch(console.error);
