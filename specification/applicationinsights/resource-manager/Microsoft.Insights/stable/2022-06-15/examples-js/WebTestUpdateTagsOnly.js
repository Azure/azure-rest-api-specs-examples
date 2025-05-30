const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates the tags associated with an Application Insights web test.
 *
 * @summary Updates the tags associated with an Application Insights web test.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestUpdateTagsOnly.json
 */
async function webTestUpdateTags() {
  const subscriptionId = process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const webTestName = "my-webtest-my-component";
  const webTestTags = {
    tags: {
      color: "AzureBlue",
      customField01: "This is a random value",
      systemType: "A08",
      "hiddenLink:/subscriptions/subid/resourceGroups/myResourceGroup/providers/MicrosoftInsights/components/myComponent":
        "Resource",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.webTests.updateTags(resourceGroupName, webTestName, webTestTags);
  console.log(result);
}
