const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Managed Environment if it does not have any container apps.
 *
 * @summary Delete a Managed Environment if it does not have any container apps.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/ManagedEnvironments_Delete.json
 */
async function deleteEnvironmentByName() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "examplerg";
  const name = "examplekenv";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironments.beginDeleteAndWait(resourceGroupName, name);
  console.log(result);
}

deleteEnvironmentByName().catch(console.error);
