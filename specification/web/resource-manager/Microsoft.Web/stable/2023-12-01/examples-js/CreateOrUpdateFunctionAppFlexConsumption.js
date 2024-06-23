const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Creates a new web, mobile, or API app in an existing resource group, or updates an existing app.
 *
 * @summary Description for Creates a new web, mobile, or API app in an existing resource group, or updates an existing app.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/CreateOrUpdateFunctionAppFlexConsumption.json
 */
async function createOrUpdateFlexConsumptionFunctionApp() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const name = "sitef6141";
  const siteEnvelope = {
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
        instanceMemoryMB: 2048,
        maximumInstanceCount: 100,
      },
    },
    kind: "functionapp,linux",
    location: "East US",
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
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.beginCreateOrUpdateAndWait(
    resourceGroupName,
    name,
    siteEnvelope,
  );
  console.log(result);
}
