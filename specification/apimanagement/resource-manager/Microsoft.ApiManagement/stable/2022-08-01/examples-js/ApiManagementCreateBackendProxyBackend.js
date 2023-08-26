const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates a backend.
 *
 * @summary Creates or Updates a backend.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateBackendProxyBackend.json
 */
async function apiManagementCreateBackendProxyBackend() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const backendId = "proxybackend";
  const parameters = {
    description: "description5308",
    credentials: {
      authorization: { parameter: "opensesma", scheme: "Basic" },
      header: { xMy1: ["val1", "val2"] },
      query: { sv: ["xx", "bb", "cc"] },
    },
    proxy: {
      password: "<password>",
      url: "http://192.168.1.1:8080",
      username: "Contoso\\admin",
    },
    tls: { validateCertificateChain: true, validateCertificateName: true },
    url: "https://backendname2644/",
    protocol: "http",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.backend.createOrUpdate(
    resourceGroupName,
    serviceName,
    backendId,
    parameters
  );
  console.log(result);
}
