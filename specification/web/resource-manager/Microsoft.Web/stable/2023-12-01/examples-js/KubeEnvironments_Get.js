const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get the properties of a Kubernetes Environment.
 *
 * @summary Description for Get the properties of a Kubernetes Environment.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/KubeEnvironments_Get.json
 */
async function getKubeEnvironmentsByName() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "examplerg";
  const name = "jlaw-demo1";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.kubeEnvironments.get(resourceGroupName, name);
  console.log(result);
}
