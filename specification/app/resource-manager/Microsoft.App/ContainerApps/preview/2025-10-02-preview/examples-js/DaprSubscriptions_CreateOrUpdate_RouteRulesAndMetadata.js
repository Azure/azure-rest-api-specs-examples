const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a Dapr subscription in a Managed Environment.
 *
 * @summary creates or updates a Dapr subscription in a Managed Environment.
 * x-ms-original-file: 2025-10-02-preview/DaprSubscriptions_CreateOrUpdate_RouteRulesAndMetadata.json
 */
async function createOrUpdateDaprSubscriptionWithRouteRulesAndMetadata() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.daprSubscriptions.createOrUpdate(
    "examplerg",
    "myenvironment",
    "mysubscription",
    {
      metadata: { foo: "bar", hello: "world" },
      pubsubName: "mypubsubcomponent",
      routes: {
        default: "/products",
        rules: [
          { path: "/widgets", match: "event.type == 'widget'" },
          { path: "/gadgets", match: "event.type == 'gadget'" },
        ],
      },
      topic: "inventory",
    },
  );
  console.log(result);
}
