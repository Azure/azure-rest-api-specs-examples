const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create or update a Container App.
 *
 * @summary Create or update a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2025-02-02-preview/examples/ContainerApps_Kind_CreateOrUpdate.json
 */
async function createOrUpdateAppKind() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const containerAppName = "testcontainerAppKind";
  const containerAppEnvelope = {
    configuration: {
      activeRevisionsMode: "Single",
      ingress: { allowInsecure: true, external: true, targetPort: 80 },
    },
    kind: "workflowapp",
    location: "East Us",
    managedBy:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppKind",
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
      scale: {
        cooldownPeriod: 350,
        maxReplicas: 30,
        minReplicas: 1,
        pollingInterval: 35,
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.beginCreateOrUpdateAndWait(
    resourceGroupName,
    containerAppName,
    containerAppEnvelope,
  );
  console.log(result);
}
