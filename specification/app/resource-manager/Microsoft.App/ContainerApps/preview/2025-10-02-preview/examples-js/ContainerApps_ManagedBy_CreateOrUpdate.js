const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Container App.
 *
 * @summary create or update a Container App.
 * x-ms-original-file: 2025-10-02-preview/ContainerApps_ManagedBy_CreateOrUpdate.json
 */
async function createOrUpdateManagedByApp() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.createOrUpdate("rg", "testcontainerAppManagedBy", {
    location: "East US",
    managedBy:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.AppPlatform/Spring/springapp",
    configuration: {
      ingress: {
        exposedPort: 4000,
        external: true,
        targetPort: 3000,
        traffic: [{ revisionName: "testcontainerAppManagedBy-ab1234", weight: 100 }],
        transport: "tcp",
      },
    },
    environmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
    template: {
      containers: [
        {
          name: "testcontainerAppManagedBy",
          image: "repo/testcontainerAppManagedBy:v1",
          probes: [
            {
              type: "Liveness",
              initialDelaySeconds: 3,
              periodSeconds: 3,
              tcpSocket: { port: 8080 },
            },
          ],
        },
      ],
      scale: {
        cooldownPeriod: 350,
        maxReplicas: 5,
        minReplicas: 1,
        pollingInterval: 35,
        rules: [{ name: "tcpscalingrule", tcp: { metadata: { concurrentConnections: "50" } } }],
      },
    },
  });
  console.log(result);
}
