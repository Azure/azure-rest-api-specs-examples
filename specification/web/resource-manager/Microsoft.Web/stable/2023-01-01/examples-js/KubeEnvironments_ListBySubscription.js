const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get all Kubernetes Environments for a subscription.
 *
 * @summary Description for Get all Kubernetes Environments for a subscription.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/KubeEnvironments_ListBySubscription.json
 */
async function listKubeEnvironmentsBySubscription() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.kubeEnvironments.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
