const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Delete a Kubernetes Environment.
 *
 * @summary Description for Delete a Kubernetes Environment.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/KubeEnvironments_Delete.json
 */
async function deleteKubeEnvironmentByName() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "examplerg";
  const name = "examplekenv";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.kubeEnvironments.beginDeleteAndWait(resourceGroupName, name);
  console.log(result);
}

deleteKubeEnvironmentByName().catch(console.error);
