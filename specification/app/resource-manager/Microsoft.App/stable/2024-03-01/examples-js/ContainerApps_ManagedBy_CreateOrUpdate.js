const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Container App.
 *
 * @summary Create or update a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ContainerApps_ManagedBy_CreateOrUpdate.json
 */
async function createOrUpdateManagedByApp() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const containerAppName = "testcontainerappmanagedby";
  const containerAppEnvelope = {
    configuration: {
      ingress: {
        exposedPort: 4000,
        external: true,
        targetPort: 3000,
        traffic: [{ revisionName: "testcontainerappmanagedby-ab1234", weight: 100 }],
        transport: "tcp",
      },
    },
    environmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
    location: "East US",
    managedBy:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.AppPlatform/Spring/springapp",
    template: {
      containers: [
        {
          name: "testcontainerappmanagedby",
          image: "repo/testcontainerappmanagedby:v1",
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
        maxReplicas: 5,
        minReplicas: 1,
        rules: [
          {
            name: "tcpscalingrule",
            tcp: { metadata: { concurrentConnections: "50" } },
          },
        ],
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
