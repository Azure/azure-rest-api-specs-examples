const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to description for Creates a new web, mobile, or API app in an existing resource group, or updates an existing app.
 *
 * @summary description for Creates a new web, mobile, or API app in an existing resource group, or updates an existing app.
 * x-ms-original-file: 2025-05-01/CreateOrUpdateFunctionAppFlexConsumptionWithDetails.json
 */
async function createOrUpdateFlexConsumptionFunctionAppWithDetails() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.createOrUpdate("testrg123", "sitef6141", {
    kind: "functionapp,linux",
    location: "East US",
    functionAppConfig: {
      deployment: {
        storage: {
          type: "blobContainer",
          authentication: {
            type: "StorageAccountConnectionString",
            storageAccountConnectionStringName: "TheAppSettingName",
          },
          value: "https://storageAccountName.blob.core.windows.net/containername",
        },
      },
      runtime: { name: "python", version: "3.11" },
      scaleAndConcurrency: {
        alwaysReady: [{ name: "http", instanceCount: 2 }],
        instanceMemoryMB: 2048,
        maximumInstanceCount: 100,
        triggers: { http: { perInstanceConcurrency: 16 } },
      },
      siteUpdateStrategy: { type: "RollingUpdate" },
    },
    siteConfig: {
      appSettings: [
        {
          name: "AzureWebJobsStorage",
          value:
            "DefaultEndpointsProtocol=https;AccountName=StorageAccountName;AccountKey=Sanitized;EndpointSuffix=core.windows.net",
        },
        {
          name: "APPLICATIONINSIGHTS_CONNECTION_STRING",
          value: "InstrumentationKey=Sanitized;IngestionEndpoint=Sanitized;LiveEndpoint=Sanitized",
        },
      ],
    },
  });
  console.log(result);
}
