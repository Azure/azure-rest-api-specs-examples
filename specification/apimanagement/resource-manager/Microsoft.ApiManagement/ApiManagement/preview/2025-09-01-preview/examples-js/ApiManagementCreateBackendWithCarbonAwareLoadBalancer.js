const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or Updates a backend.
 *
 * @summary creates or Updates a backend.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateBackendWithCarbonAwareLoadBalancer.json
 */
async function apiManagementCreateBackendWithCarbonAwareLoadBalancer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.backend.createOrUpdate(
    "rg1",
    "apimService1",
    "carbonawareloadbalancerbackend",
    {
      typePropertiesType: "Pool",
      pool: {
        services: [
          {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/sustainable-backend-europe-north",
            preferredCarbonEmission: "Medium",
            priority: 1,
            weight: 1,
          },
          {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/proxybackend",
            priority: 1,
            weight: 1,
          },
        ],
      },
    },
  );
  console.log(result);
}
