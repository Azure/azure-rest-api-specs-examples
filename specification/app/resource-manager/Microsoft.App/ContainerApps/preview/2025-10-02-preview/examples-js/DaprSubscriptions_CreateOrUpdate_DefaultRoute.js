const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a Dapr subscription in a Managed Environment.
 *
 * @summary creates or updates a Dapr subscription in a Managed Environment.
 * x-ms-original-file: 2025-10-02-preview/DaprSubscriptions_CreateOrUpdate_DefaultRoute.json
 */
async function createOrUpdateDaprSubscriptionWithDefaultRouteOnly() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.daprSubscriptions.createOrUpdate(
    "examplerg",
    "myenvironment",
    "mysubscription",
    { pubsubName: "mypubsubcomponent", routes: { default: "/products" }, topic: "inventory" },
  );
  console.log(result);
}
