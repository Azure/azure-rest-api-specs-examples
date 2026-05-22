const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Container App.
 *
 * @summary create or update a Container App.
 * x-ms-original-file: 2025-10-02-preview/ContainerApps_Kind_WorkflowApp_CreateOrUpdate.json
 */
async function createOrUpdateWorkflowAppKind() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.createOrUpdate("rg", "testcontainerAppKind", {
    kind: "workflowapp",
    location: "East Us",
    managedBy:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppKind",
    configuration: {
      activeRevisionsMode: "Single",
      ingress: { allowInsecure: false, external: true, targetPort: 443 },
    },
    managedEnvironmentId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3",
    template: {
      containers: [
        {
          name: "logicapps-container",
          image: "default/logicapps-base:latest",
          resources: { cpu: 1, memory: "2.0Gi" },
        },
      ],
      scale: { cooldownPeriod: 350, maxReplicas: 30, minReplicas: 1, pollingInterval: 35 },
    },
  });
  console.log(result);
}
