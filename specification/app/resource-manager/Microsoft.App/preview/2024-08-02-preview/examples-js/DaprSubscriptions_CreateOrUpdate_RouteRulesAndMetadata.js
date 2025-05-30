const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Dapr subscription in a Managed Environment.
 *
 * @summary Creates or updates a Dapr subscription in a Managed Environment.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/DaprSubscriptions_CreateOrUpdate_RouteRulesAndMetadata.json
 */
async function createOrUpdateDaprSubscriptionWithRouteRulesAndMetadata() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "examplerg";
  const environmentName = "myenvironment";
  const name = "mysubscription";
  const daprSubscriptionEnvelope = {
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
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.daprSubscriptions.createOrUpdate(
    resourceGroupName,
    environmentName,
    name,
    daprSubscriptionEnvelope,
  );
  console.log(result);
}
