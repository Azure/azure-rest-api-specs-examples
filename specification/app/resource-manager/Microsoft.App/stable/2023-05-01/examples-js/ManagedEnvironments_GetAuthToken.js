const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks if resource name is available.
 *
 * @summary Checks if resource name is available.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2023-05-01/examples/ManagedEnvironments_GetAuthToken.json
 */
async function getManagedEnvironmentAuthToken() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const environmentName = "testenv";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironments.getAuthToken(resourceGroupName, environmentName);
  console.log(result);
}
