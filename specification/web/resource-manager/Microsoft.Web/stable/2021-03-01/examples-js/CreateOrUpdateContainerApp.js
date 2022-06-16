const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Create or update a Container App.
 *
 * @summary Description for Create or update a Container App.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/CreateOrUpdateContainerApp.json
 */
async function createOrUpdateContainerApp() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testcontainerApp0";
  const containerAppEnvelope = {
    configuration: { ingress: { external: true, targetPort: 3000 } },
    kind: "containerApp",
    kubeEnvironmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.Web/kubeEnvironments/demokube",
    location: "East US",
    template: {
      containers: [{ name: "testcontainerApp0", image: "repo/testcontainerApp0:v1" }],
      dapr: { appPort: 3000, enabled: true },
      scale: {
        maxReplicas: 5,
        minReplicas: 1,
        rules: [
          {
            name: "httpscalingrule",
            custom: { type: "http", metadata: { concurrentRequests: "50" } },
          },
        ],
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.containerApps.beginCreateOrUpdateAndWait(
    resourceGroupName,
    name,
    containerAppEnvelope
  );
  console.log(result);
}

createOrUpdateContainerApp().catch(console.error);
