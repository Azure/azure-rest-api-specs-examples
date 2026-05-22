const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates an Application Insights web test definition.
 *
 * @summary creates or updates an Application Insights web test definition.
 * x-ms-original-file: 2022-06-15/WebTestCreate.json
 */
async function webTestCreate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.webTests.createOrUpdate(
    "my-resource-group",
    "my-webtest-my-component",
    {
      kind: "ping",
      location: "South Central US",
      configuration: {
        webTest:
          '<WebTest Name="my-webtest" Id="678ddf96-1ab8-44c8-9274-123456789abc" Enabled="True" CssProjectStructure="" CssIteration="" Timeout="120" WorkItemIds="" xmlns="http://microsoft.com/schemas/VisualStudio/TeamTest/2010" Description="" CredentialUserName="" CredentialPassword="" PreAuthenticate="True" Proxy="default" StopOnError="False" RecordedResultFile="" ResultsLocale="" ><Items><Request Method="GET" Guid="a4162485-9114-fcfc-e086-123456789abc" Version="1.1" Url="http://my-component.azurewebsites.net" ThinkTime="0" Timeout="120" ParseDependentRequests="True" FollowRedirects="True" RecordResult="True" Cache="False" ResponseTimeGoal="0" Encoding="utf-8" ExpectedHttpStatusCode="200" ExpectedResponseUrl="" ReportingName="" IgnoreHttpStatusCode="False" /></Items></WebTest>',
      },
      description: "Ping web test alert for mytestwebapp",
      enabled: true,
      frequency: 900,
      webTestKind: "ping",
      locations: [{ location: "us-fl-mia-edge" }],
      webTestName: "my-webtest-my-component",
      retryEnabled: true,
      syntheticMonitorId: "my-webtest-my-component",
      timeout: 120,
    },
  );
  console.log(result);
}
