const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Container App.
 *
 * @summary Create or update a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/ContainerApps_TcpApp_CreateOrUpdate.json
 */
async function createOrUpdateTcpApp() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const containerAppName = "testcontainerAppTcp";
  const containerAppEnvelope = {
    configuration: {
      ingress: {
        exposedPort: 4000,
        external: true,
        targetPort: 3000,
        traffic: [{ revisionName: "testcontainerAppTcp-ab1234", weight: 100 }],
        transport: "tcp",
      },
    },
    environmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
    location: "East US",
    template: {
      containers: [
        {
          name: "testcontainerAppTcp",
          image: "repo/testcontainerAppTcp:v1",
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
