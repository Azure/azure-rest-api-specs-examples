const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or Updates a backend.
 *
 * @summary creates or Updates a backend.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateWorkspaceBackendServiceFabric.json
 */
async function apiManagementCreateWorkspaceBackendServiceFabric() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceBackend.createOrUpdate(
    "rg1",
    "apimService1",
    "wks1",
    "sfbackend",
    {
      description: "Service Fabric Test App 1",
      properties: {
        serviceFabricCluster: {
          clientCertificateId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/workspaces/wks1/certificates/cert1",
          managementEndpoints: ["https://somecluster.com"],
          maxPartitionResolutionRetries: 5,
          serverX509Names: [
            {
              name: "ServerCommonName1",
              issuerCertificateThumbprint: "IssuerCertificateThumbprint1",
            },
          ],
        },
      },
      url: "fabric:/mytestapp/mytestservice",
      protocol: "http",
    },
  );
  console.log(result);
}
