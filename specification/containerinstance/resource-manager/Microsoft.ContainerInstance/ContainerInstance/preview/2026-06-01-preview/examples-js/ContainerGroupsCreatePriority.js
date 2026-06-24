const { ContainerInstanceManagementClient } = require("@azure/arm-containerinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update container groups with specified configurations.
 *
 * @summary create or update container groups with specified configurations.
 * x-ms-original-file: 2026-06-01-preview/ContainerGroupsCreatePriority.json
 */
async function containerGroupsCreateWithPriority() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerInstanceManagementClient(credential, subscriptionId);
  const result = await client.containerGroups.createOrUpdate("demo", "demo1", {
    location: "eastus",
    containers: [
      {
        name: "test-container-001",
        command: ["/bin/sh", "-c", "sleep 10"],
        image: "alpine:latest",
        resources: { requests: { cpu: 1, memoryInGB: 1 } },
      },
    ],
    osType: "Linux",
    priority: "Spot",
    restartPolicy: "Never",
    sku: "Standard",
  });
  console.log(result);
}
