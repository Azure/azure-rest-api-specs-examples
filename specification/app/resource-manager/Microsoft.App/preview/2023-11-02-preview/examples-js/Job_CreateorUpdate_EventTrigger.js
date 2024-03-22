const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or Update a Container Apps Job.
 *
 * @summary Create or Update a Container Apps Job.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Job_CreateorUpdate_EventTrigger.json
 */
async function createOrUpdateContainerAppsJobWithEventDrivenTrigger() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const jobName = "testcontainerAppsJob0";
  const jobEnvelope = {
    configuration: {
      eventTriggerConfig: {
        parallelism: 4,
        replicaCompletionCount: 1,
        scale: {
          maxExecutions: 5,
          minExecutions: 1,
          pollingInterval: 40,
          rules: [
            {
              name: "servicebuscalingrule",
              type: "azure-servicebus",
              metadata: { topicName: "my-topic" },
            },
          ],
        },
      },
      replicaRetryLimit: 10,
      replicaTimeout: 10,
      triggerType: "Event",
    },
    environmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
    location: "East US",
    template: {
      containers: [
        {
          name: "testcontainerAppsJob0",
          image: "repo/testcontainerAppsJob0:v1",
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
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.jobs.beginCreateOrUpdateAndWait(
    resourceGroupName,
    jobName,
    jobEnvelope,
  );
  console.log(result);
}
