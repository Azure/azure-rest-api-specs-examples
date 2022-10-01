const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Activates a revision for a Container App
 *
 * @summary Activates a revision for a Container App
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/Revisions_Activate.json
 */
async function activateContainerAppRevision() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const containerAppName = "testcontainerApp0";
  const name = "testcontainerApp0-pjxhsye";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsRevisions.activateRevision(
    resourceGroupName,
    containerAppName,
    name
  );
  console.log(result);
}

activateContainerAppRevision().catch(console.error);
