const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Container App.
 *
 * @summary Create or update a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ContainerApps_CreateOrUpdate.json
 */
async function createOrUpdateContainerApp() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const containerAppName = "testcontainerApp0";
  const containerAppEnvelope = {
    configuration: {
      dapr: { appPort: 3000, appProtocol: "http", enabled: true },
      ingress: {
        customDomains: [
          {
            name: "www.my-name.com",
            bindingType: "SniEnabled",
            certificateId:
              "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-name-dot-com",
          },
          {
            name: "www.my-other-name.com",
            bindingType: "SniEnabled",
            certificateId:
              "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube/certificates/my-certificate-for-my-other-name-dot-com",
          },
        ],
        external: true,
        targetPort: 3000,
        traffic: [
          {
            label: "production",
            revisionName: "testcontainerApp0-ab1234",
            weight: 100,
          },
        ],
      },
    },
    location: "East US",
    managedEnvironmentId:
      "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/demokube",
    template: {
      containers: [
        {
          name: "testcontainerApp0",
          image: "repo/testcontainerApp0:v1",
          probes: [
            {
              type: "Liveness",
              httpGet: {
                path: "/health",
                httpHeaders: [{ name: "Custom-Header", value: "Awesome" }],
                port: 8080,
              },
              initialDelaySeconds: 3,
              periodSeconds: 3,
            },
          ],
        },
      ],
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
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.beginCreateOrUpdateAndWait(
    resourceGroupName,
    containerAppName,
    containerAppEnvelope
  );
  console.log(result);
}

createOrUpdateContainerApp().catch(console.error);
