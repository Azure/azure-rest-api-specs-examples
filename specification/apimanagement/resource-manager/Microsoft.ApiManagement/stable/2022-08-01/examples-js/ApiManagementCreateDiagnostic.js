const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new Diagnostic or updates an existing one.
 *
 * @summary Creates a new Diagnostic or updates an existing one.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateDiagnostic.json
 */
async function apiManagementCreateDiagnostic() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const diagnosticId = "applicationinsights";
  const parameters = {
    alwaysLog: "allErrors",
    backend: {
      response: { body: { bytes: 512 }, headers: ["Content-type"] },
      request: { body: { bytes: 512 }, headers: ["Content-type"] },
    },
    frontend: {
      response: { body: { bytes: 512 }, headers: ["Content-type"] },
      request: { body: { bytes: 512 }, headers: ["Content-type"] },
    },
    loggerId: "/loggers/azuremonitor",
    sampling: { percentage: 50, samplingType: "fixed" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.diagnostic.createOrUpdate(
    resourceGroupName,
    serviceName,
    diagnosticId,
    parameters
  );
  console.log(result);
}
