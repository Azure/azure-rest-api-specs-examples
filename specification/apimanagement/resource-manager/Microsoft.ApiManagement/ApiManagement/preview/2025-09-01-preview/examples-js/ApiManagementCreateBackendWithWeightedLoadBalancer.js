const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or Updates a backend.
 *
 * @summary creates or Updates a backend.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateBackendWithWeightedLoadBalancer.json
 */
async function apiManagementCreateBackendWithWeightedLoadBalancer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.backend.createOrUpdate(
    "rg1",
    "apimService1",
    "weighted-load-balancer",
    {
      typePropertiesType: "Pool",
      pool: {
        services: [
          {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/backend-1",
            weight: 75,
          },
          {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/backend-2",
            weight: 25,
          },
        ],
      },
    },
  );
  console.log(result);
}
