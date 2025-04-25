const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or Updates a backend.
 *
 * @summary Creates or Updates a backend.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateWorkspaceBackendServiceFabric.json
 */
async function apiManagementCreateWorkspaceBackendServiceFabric() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const workspaceId = "wks1";
  const backendId = "sfbackend";
  const parameters = {
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
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.workspaceBackend.createOrUpdate(
    resourceGroupName,
    serviceName,
    workspaceId,
    backendId,
    parameters,
  );
  console.log(result);
}
