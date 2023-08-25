const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or Updates a backend.
 *
 * @summary Creates or Updates a backend.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateBackendServiceFabric.json
 */
async function apiManagementCreateBackendServiceFabric() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const backendId = "sfbackend";
  const parameters = {
    description: "Service Fabric Test App 1",
    properties: {
      serviceFabricCluster: {
        clientCertificateId:
          "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/certificates/cert1",
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
  const result = await client.backend.createOrUpdate(
    resourceGroupName,
    serviceName,
    backendId,
    parameters
  );
  console.log(result);
}
