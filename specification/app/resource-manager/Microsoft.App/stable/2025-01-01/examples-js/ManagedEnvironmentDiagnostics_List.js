const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Get the list of diagnostics for a Managed Environment used to host container apps.
 *
 * @summary Get the list of diagnostics for a Managed Environment used to host container apps.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/ManagedEnvironmentDiagnostics_List.json
 */
async function getTheListOfAvailableDiagnosticDataForAManagedEnvironments() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "f07f3711-b45e-40fe-a941-4e6d93f851e6";
  const resourceGroupName =
    process.env["APPCONTAINERS_RESOURCE_GROUP"] || "mikono-workerapp-test-rg";
  const environmentName = "mikonokubeenv";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironmentDiagnostics.listDetectors(
    resourceGroupName,
    environmentName,
  );
  console.log(result);
}
