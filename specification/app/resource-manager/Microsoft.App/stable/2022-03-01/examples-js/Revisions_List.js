const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the Revisions for a given Container App.
 *
 * @summary Get the Revisions for a given Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/Revisions_List.json
 */
async function listContainerAppRevisions() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const containerAppName = "testcontainerApp0";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.containerAppsRevisions.listRevisions(
    resourceGroupName,
    containerAppName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listContainerAppRevisions().catch(console.error);
