const { ManagedDevOpsInfrastructure } = require("@azure/arm-devopsinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Pool
 *
 * @summary Create a Pool
 * x-ms-original-file: specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/CreateOrUpdatePool.json
 */
async function poolsCreateOrUpdate() {
  const subscriptionId =
    process.env["DEVOPSINFRASTRUCTURE_SUBSCRIPTION_ID"] || "a2e95d27-c161-4b61-bda4-11512c14c2c2";
  const resourceGroupName = process.env["DEVOPSINFRASTRUCTURE_RESOURCE_GROUP"] || "rg";
  const poolName = "pool";
  const resource = {
    location: "eastus",
    properties: {
      agentProfile: { kind: "Stateless" },
      devCenterProjectResourceId:
        "/subscriptions/222e81d0-cf38-4dab-baa5-289bf16baaa4/resourceGroups/rg-1es-devcenter/providers/Microsoft.DevCenter/projects/1ES",
      fabricProfile: {
        images: [
          {
            resourceId: "/MicrosoftWindowsServer/WindowsServer/2019-Datacenter/latest",
          },
        ],
        kind: "Vmss",
        sku: { name: "Standard_D4ads_v5" },
      },
      maximumConcurrency: 10,
      organizationProfile: {
        kind: "AzureDevOps",
        organizations: [{ url: "https://mseng.visualstudio.com" }],
      },
      provisioningState: "Succeeded",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ManagedDevOpsInfrastructure(credential, subscriptionId);
  const result = await client.pools.beginCreateOrUpdateAndWait(
    resourceGroupName,
    poolName,
    resource,
  );
  console.log(result);
}
