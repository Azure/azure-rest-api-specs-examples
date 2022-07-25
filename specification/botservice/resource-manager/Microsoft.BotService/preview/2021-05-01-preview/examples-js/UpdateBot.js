const { AzureBotService } = require("@azure/arm-botservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a Bot Service
 *
 * @summary Updates a Bot Service
 * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/UpdateBot.json
 */
async function updateBot() {
  const subscriptionId = "subscription-id";
  const resourceGroupName = "OneResourceGroupName";
  const resourceName = "samplebotname";
  const location = "West US";
  const tags = { tag1: "value1", tag2: "value2" };
  const sku = { name: "S1" };
  const kind = "sdk";
  const etag = "etag1";
  const properties = {
    description: "The description of the bot",
    cmekKeyVaultUrl: "https://myCmekKey",
    developerAppInsightKey: "appinsightskey",
    developerAppInsightsApiKey: "appinsightsapikey",
    developerAppInsightsApplicationId: "appinsightsappid",
    disableLocalAuth: true,
    displayName: "The Name of the bot",
    endpoint: "http://mybot.coffee",
    iconUrl: "http://myicon",
    isCmekEnabled: true,
    luisAppIds: ["luisappid1", "luisappid2"],
    luisKey: "luiskey",
    msaAppId: "msaappid",
    msaAppMSIResourceId:
      "/subscriptions/foo/resourcegroups/bar/providers/microsoft.managedidentity/userassignedidentities/sampleId",
    msaAppTenantId: "msaapptenantid",
    msaAppType: "UserAssignedMSI",
    publicNetworkAccess: "Enabled",
    schemaTransformationVersion: "1.0",
  };
  const options = {
    location,
    tags,
    sku,
    kind,
    etag,
    properties,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureBotService(credential, subscriptionId);
  const result = await client.bots.update(resourceGroupName, resourceName, options);
  console.log(result);
}

updateBot().catch(console.error);
