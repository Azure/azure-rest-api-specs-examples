const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches a session pool using JSON merge patch
 *
 * @summary Patches a session pool using JSON merge patch
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/SessionPools_Patch.json
 */
async function updateSessionPool() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const sessionPoolName = "testsessionpool";
  const sessionPoolEnvelope = {
    customContainerTemplate: {
      containers: [
        {
          name: "testinitcontainer",
          args: ["-c", "while true; do echo hello; sleep 10;done"],
          command: ["/bin/sh"],
          image: "repo/testcontainer:v4",
          resources: { cpu: 0.25, memory: "0.5Gi" },
        },
      ],
      ingress: { targetPort: 80 },
    },
    dynamicPoolConfiguration: {
      cooldownPeriodInSeconds: 600,
      executionType: "Timed",
    },
    scaleConfiguration: {
      maxConcurrentSessions: 500,
      readySessionInstances: 100,
    },
    sessionNetworkConfiguration: { status: "EgressEnabled" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsSessionPools.beginUpdateAndWait(
    resourceGroupName,
    sessionPoolName,
    sessionPoolEnvelope,
  );
  console.log(result);
}
