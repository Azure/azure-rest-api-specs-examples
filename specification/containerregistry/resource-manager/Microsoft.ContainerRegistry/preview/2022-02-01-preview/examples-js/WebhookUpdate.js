const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a webhook with the specified parameters.
 *
 * @summary Updates a webhook with the specified parameters.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2022-02-01-preview/examples/WebhookUpdate.json
 */
async function webhookUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const registryName = "myRegistry";
  const webhookName = "myWebhook";
  const webhookUpdateParameters = {
    actions: ["push"],
    customHeaders: {
      authorization: "",
    },
    scope: "myRepository",
    serviceUri: "http://myservice.com",
    status: "enabled",
    tags: { key: "value" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.webhooks.beginUpdateAndWait(
    resourceGroupName,
    registryName,
    webhookName,
    webhookUpdateParameters
  );
  console.log(result);
}

webhookUpdate().catch(console.error);
