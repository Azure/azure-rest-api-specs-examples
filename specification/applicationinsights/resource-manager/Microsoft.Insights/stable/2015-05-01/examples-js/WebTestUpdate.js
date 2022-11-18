const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Application Insights web test definition.
 *
 * @summary Creates or updates an Application Insights web test definition.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WebTestUpdate.json
 */
async function webTestUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const webTestName = "my-webtest-my-component";
  const webTestDefinition = {
    configuration: {
      webTest:
        '<WebTest Name="my-webtest" Id="678ddf96-1ab8-44c8-9274-123456789abc" Enabled="True" CssProjectStructure="" CssIteration="" Timeout="30" WorkItemIds="" xmlns="http://microsoft.com/schemas/VisualStudio/TeamTest/2010" Description="" CredentialUserName="" CredentialPassword="" PreAuthenticate="True" Proxy="default" StopOnError="False" RecordedResultFile="" ResultsLocale="" ><Items><Request Method="GET" Guid="a4162485-9114-fcfc-e086-123456789abc" Version="1.1" Url="http://my-component.azurewebsites.net" ThinkTime="0" Timeout="30" ParseDependentRequests="True" FollowRedirects="True" RecordResult="True" Cache="False" ResponseTimeGoal="0" Encoding="utf-8" ExpectedHttpStatusCode="200" ExpectedResponseUrl="" ReportingName="" IgnoreHttpStatusCode="False" /></Items></WebTest>',
    },
    frequency: 600,
    webTestKind: "ping",
    locations: [{ location: "us-fl-mia-edge" }, { location: "apac-hk-hkn-azr" }],
    webTestName: "my-webtest-my-component",
    syntheticMonitorId: "my-webtest-my-component",
    timeout: 30,
    kind: "ping",
    location: "South Central US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.webTests.createOrUpdate(
    resourceGroupName,
    webTestName,
    webTestDefinition
  );
  console.log(result);
}

webTestUpdate().catch(console.error);
