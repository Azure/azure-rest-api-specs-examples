const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Restarts a revision for a Container App
 *
 * @summary Restarts a revision for a Container App
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Revisions_Restart.json
 */
async function restartContainerAppRevision() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const containerAppName = "testStaticSite0";
  const revisionName = "testcontainerApp0-pjxhsye";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsRevisions.restartRevision(
    resourceGroupName,
    containerAppName,
    revisionName,
  );
  console.log(result);
}
