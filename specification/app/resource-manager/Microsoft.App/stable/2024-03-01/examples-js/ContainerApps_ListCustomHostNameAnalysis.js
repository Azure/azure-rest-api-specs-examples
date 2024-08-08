const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Analyzes a custom hostname for a Container App
 *
 * @summary Analyzes a custom hostname for a Container App
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/ContainerApps_ListCustomHostNameAnalysis.json
 */
async function analyzeCustomHostname() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const containerAppName = "testcontainerapp0";
  const customHostname = "my.name.corp";
  const options = {
    customHostname,
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.listCustomHostNameAnalysis(
    resourceGroupName,
    containerAppName,
    options,
  );
  console.log(result);
}
