const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or Update a Container Apps Job.
 *
 * @summary create or Update a Container Apps Job.
 * x-ms-original-file: 2025-10-02-preview/Job_CreateorUpdate_ConnectedEnvironment.json
 */
async function createOrUpdateContainerAppsJobOnAConnectedEnvironment() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate("rg", "testcontainerAppsJob0", {
    extendedLocation: {
      name: "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.ExtendedLocation/customLocations/testcustomlocation",
      type: "CustomLocation",
    },
    location: "East US",
    configuration: {
      manualTriggerConfig: { parallelism: 4, replicaCompletionCount: 1 },
      replicaRetryLimit: 10,
      replicaTimeout: 10,
      triggerType: "Manual",
    },
    environmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/connectedEnvironments/demokube",
    template: {
      containers: [
        {
          name: "testcontainerAppsJob0",
          image: "repo/testcontainerAppsJob0:v1",
          probes: [
            {
              type: "Liveness",
              httpGet: {
                path: "/health",
                httpHeaders: [{ name: "Custom-Header", value: "Awesome" }],
                port: 8080,
              },
              initialDelaySeconds: 5,
              periodSeconds: 3,
            },
          ],
        },
      ],
      initContainers: [
        {
          name: "testinitcontainerAppsJob0",
          args: ["-c", "while true; do echo hello; sleep 10;done"],
          command: ["/bin/sh"],
          image: "repo/testcontainerAppsJob0:v4",
          resources: { cpu: 0.2, memory: "100Mi" },
        },
      ],
    },
  });
  console.log(result);
}
