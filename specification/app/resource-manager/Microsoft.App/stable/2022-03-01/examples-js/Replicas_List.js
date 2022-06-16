const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List replicas for a Container App Revision.
 *
 * @summary List replicas for a Container App Revision.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/Replicas_List.json
 */
async function listContainerAppReplicas() {
  const subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const resourceGroupName = "workerapps-rg-xj";
  const containerAppName = "myapp";
  const revisionName = "myapp--0wlqy09";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsRevisionReplicas.listReplicas(
    resourceGroupName,
    containerAppName,
    revisionName
  );
  console.log(result);
}

listContainerAppReplicas().catch(console.error);
