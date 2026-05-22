const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a session pool with the given properties.
 *
 * @summary create or update a session pool with the given properties.
 * x-ms-original-file: 2025-10-02-preview/SessionPools_McpServer_CreateOrUpdate.json
 */
async function createOrUpdateSessionPoolWithMCPServer() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsSessionPools.createOrUpdate("rg", "testsessionpool", {
    location: "East US",
    containerType: "Shell",
    dynamicPoolConfiguration: {
      lifecycleConfiguration: { cooldownPeriodInSeconds: 600, lifecycleType: "Timed" },
    },
    mcpServerSettings: { isMcpServerApiKeyDisabled: false, isMcpServerEnabled: true },
    poolManagementType: "Dynamic",
    scaleConfiguration: { maxConcurrentSessions: 50 },
    sessionNetworkConfiguration: { status: "EgressEnabled" },
  });
  console.log(result);
}
