const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create an API Key of an Application Insights component.
 *
 * @summary Create an API Key of an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/APIKeysCreate.json
 */
async function apiKeyCreate() {
  const subscriptionId = process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const resourceName = "my-component";
  const aPIKeyProperties = {
    name: "test2",
    linkedReadProperties: [
      "/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/api",
      "/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/agentconfig",
    ],
    linkedWriteProperties: [
      "/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.Insights/components/my-component/annotations",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.aPIKeys.create(resourceGroupName, resourceName, aPIKeyProperties);
  console.log(result);
}
