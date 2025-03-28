const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates an Application Insights web test definition.
 *
 * @summary Creates or updates an Application Insights web test definition.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-06-15/examples/WebTestCreateStandard.json
 */
async function webTestCreateStandard() {
  const subscriptionId = process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const webTestName = "my-webtest-my-component";
  const webTestDefinition = {
    description: "Ping web test alert for mytestwebapp",
    enabled: true,
    frequency: 900,
    webTestKind: "standard",
    locations: [{ location: "us-fl-mia-edge" }],
    webTestName: "my-webtest-my-component",
    request: {
      headers: [
        { headerFieldName: "Content-Language", headerFieldValue: "de-DE" },
        { headerFieldName: "Accept-Language", headerFieldValue: "de-DE" },
      ],
      httpVerb: "POST",
      requestBody: "SGVsbG8gd29ybGQ=",
      requestUrl: "https://bing.com",
    },
    retryEnabled: true,
    syntheticMonitorId: "my-webtest-my-component",
    timeout: 120,
    validationRules: { sSLCertRemainingLifetimeCheck: 100, sSLCheck: true },
    location: "South Central US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.webTests.createOrUpdate(
    resourceGroupName,
    webTestName,
    webTestDefinition,
  );
  console.log(result);
}
