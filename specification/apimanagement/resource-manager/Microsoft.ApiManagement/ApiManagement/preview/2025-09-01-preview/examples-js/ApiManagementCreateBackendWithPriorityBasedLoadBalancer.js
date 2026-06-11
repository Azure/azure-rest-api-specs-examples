const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or Updates a backend.
 *
 * @summary creates or Updates a backend.
 * x-ms-original-file: 2025-09-01-preview/ApiManagementCreateBackendWithPriorityBasedLoadBalancer.json
 */
async function apiManagementCreateBackendWithPriorityBasedLoadBalancer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.backend.createOrUpdate(
    "rg1",
    "apimService1",
    "priority-based-load-balancer",
    {
      typePropertiesType: "Pool",
      pool: {
        services: [
          {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/backend-1",
            priority: 1,
          },
          {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/backend-2",
            priority: 1,
          },
          {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/backend-3",
            priority: 2,
          },
          {
            id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/backends/backend-4",
            priority: 2,
          },
        ],
      },
    },
  );
  console.log(result);
}
