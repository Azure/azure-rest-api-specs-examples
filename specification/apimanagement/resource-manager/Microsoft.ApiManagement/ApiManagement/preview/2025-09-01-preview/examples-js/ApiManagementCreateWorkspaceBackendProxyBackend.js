const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or Updates a backend.
 *
 * @summary creates or Updates a backend.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateWorkspaceBackendProxyBackend.json
 */
async function apiManagementCreateWorkspaceBackendProxyBackend() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceBackend.createOrUpdate(
    "rg1",
    "apimService1",
    "wks1",
    "proxybackend",
    {
      description: "description5308",
      credentials: {
        authorization: { parameter: "opensesma", scheme: "Basic" },
        header: { "x-my-1": ["val1", "val2"] },
        query: { sv: ["xx", "bb", "cc"] },
      },
      proxy: { password: "<password>", url: "http://192.168.1.1:8080", username: "Contoso\\admin" },
      tls: { validateCertificateChain: true, validateCertificateName: true },
      url: "https://backendname2644/",
      protocol: "http",
    },
  );
  console.log(result);
}
