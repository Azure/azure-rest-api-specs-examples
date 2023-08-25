const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing backend.
 *
 * @summary Updates an existing backend.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateBackend.json
 */
async function apiManagementUpdateBackend() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const backendId = "proxybackend";
  const ifMatch = "*";
  const parameters = {
    description: "description5308",
    tls: { validateCertificateChain: false, validateCertificateName: true },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.backend.update(
    resourceGroupName,
    serviceName,
    backendId,
    ifMatch,
    parameters
  );
  console.log(result);
}
