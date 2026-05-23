const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update a Container App.
 *
 * @summary create or update a Container App.
 * x-ms-original-file: 2025-10-02-preview/ContainerApps_Kind_FunctionApp_CreateOrUpdate.json
 */
async function createOrUpdateFunctionAppKind() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.createOrUpdate("rg", "testcontainerAppFunctionKind", {
    kind: "functionapp",
    location: "East Us",
    managedBy:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.Web/sites/testcontainerAppFunctionKind",
    configuration: {
      activeRevisionsMode: "Single",
      ingress: { allowInsecure: false, external: true, targetPort: 80 },
    },
    managedEnvironmentId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg/providers/Microsoft.App/managedEnvironments/testmanagedenv3",
    template: {
      containers: [
        {
          name: "function-app-container",
          env: [
            {
              name: "AzureWebJobsStorage",
              value:
                "DefaultEndpointsProtocol=https;AccountName=mystorageaccount;AccountKey=mykey;EndpointSuffix=core.windows.net",
            },
            { name: "FUNCTIONS_WORKER_RUNTIME", value: "dotnet" },
            { name: "WEBSITES_ENABLE_APP_SERVICE_STORAGE", value: "false" },
          ],
          image: "mcr.microsoft.com/azure-functions/dotnet:4",
          resources: { cpu: 0.5, memory: "1.0Gi" },
        },
      ],
      scale: { cooldownPeriod: 300, maxReplicas: 10, minReplicas: 0, pollingInterval: 30 },
    },
  });
  console.log(result);
}
