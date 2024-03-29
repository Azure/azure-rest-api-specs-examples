const { HealthbotClient } = require("@azure/arm-healthbot");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Azure Health Bot.
 *
 * @summary Create a new Azure Health Bot.
 * x-ms-original-file: specification/healthbot/resource-manager/Microsoft.HealthBot/stable/2021-06-10/examples/ResourceCreationPut.json
 */
async function botCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "healthbotClient";
  const botName = "samplebotname";
  const parameters = {
    identity: {
      type: "SystemAssigned, UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/subscriptionId/resourcegroups/myrg/providers/microsoftManagedidentity/userassignedidentities/myMi":
          {},
        "/subscriptions/subscriptionId/resourcegroups/myrg/providers/microsoftManagedidentity/userassignedidentities/myMi2":
          {},
      },
    },
    location: "East US",
    sku: { name: "F0" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthbotClient(credential, subscriptionId);
  const result = await client.bots.beginCreateAndWait(resourceGroupName, botName, parameters);
  console.log(result);
}

botCreate().catch(console.error);
