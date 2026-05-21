const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Application Insights web test definition.
 *
 * @summary creates or updates an Application Insights web test definition.
 * x-ms-original-file: 2022-06-15/WebTestCreateStandard.json
 */
async function webTestCreateStandard() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.webTests.createOrUpdate(
    "my-resource-group",
    "my-webtest-my-component",
    {
      location: "South Central US",
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
    },
  );
  console.log(result);
}
