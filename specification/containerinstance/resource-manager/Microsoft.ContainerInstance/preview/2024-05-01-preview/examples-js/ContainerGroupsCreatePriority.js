const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update container groups with specified configurations.
 *
 * @summary Create or update container groups with specified configurations.
 * x-ms-original-file: specification/containerinstance/resource-manager/Microsoft.ContainerInstance/preview/2024-05-01-preview/examples/ContainerGroupsCreatePriority.json
 */
async function containerGroupsCreateWithPriority() {
  const subscriptionId =
    process.env["CONTAINERINSTANCE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERINSTANCE_RESOURCE_GROUP"] || "demo";
  const containerGroupName = "demo1";
  const containerGroup = {
    containers: [
      {
        name: "test-container-001",
        command: ["/bin/sh", "-c", "sleep 10"],
        image: "alpine:latest",
        resources: { requests: { cpu: 1, memoryInGB: 1 } },
      },
    ],
    location: "eastus",
    osType: "Linux",
    priority: "Spot",
    restartPolicy: "Never",
    sku: "Standard",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    containerGroupName,
    containerGroup,
  );
  console.log(result);
}
