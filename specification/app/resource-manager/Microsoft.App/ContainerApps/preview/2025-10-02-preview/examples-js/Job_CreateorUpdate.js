const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or Update a Container Apps Job.
 *
 * @summary create or Update a Container Apps Job.
 * x-ms-original-file: 2025-10-02-preview/Job_CreateorUpdate.json
 */
async function createOrUpdateContainerAppsJob() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.jobs.createOrUpdate("rg", "testcontainerAppsJob0", {
    identity: {
      type: "SystemAssigned,UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity":
          {},
      },
    },
    location: "East US",
    configuration: {
      identitySettings: [
        {
          identity:
            "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourcegroups/rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myidentity",
          lifecycle: "All",
        },
        { identity: "system", lifecycle: "Init" },
      ],
      manualTriggerConfig: { parallelism: 4, replicaCompletionCount: 1 },
      replicaRetryLimit: 10,
      replicaTimeout: 10,
      triggerType: "Manual",
    },
    environmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
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
          volumeMounts: [
            { mountPath: "/mnt/path1", subPath: "subPath1", volumeName: "azurefile" },
            { mountPath: "/mnt/path2", subPath: "subPath2", volumeName: "nfsazurefile" },
          ],
        },
      ],
      initContainers: [
        {
          name: "testinitcontainerAppsJob0",
          args: ["-c", "while true; do echo hello; sleep 10;done"],
          command: ["/bin/sh"],
          image: "repo/testcontainerAppsJob0:v4",
          resources: { cpu: 0.2, gpu: 1, memory: "100Mi" },
        },
      ],
    },
  });
  console.log(result);
}
