const { TimeSeriesInsightsClient } = require("@azure/arm-timeseriesinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an event source under the specified environment.
 *
 * @summary Create or update an event source under the specified environment.
 * x-ms-original-file: specification/timeseriesinsights/resource-manager/Microsoft.TimeSeriesInsights/preview/2021-03-31-preview/examples/EventSourcesCreateEventHubWithCustomEnquedTime.json
 */
async function eventSourcesCreateEventHubWithCustomEnquedTime() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const environmentName = "env1";
  const eventSourceName = "es1";
  const parameters = {
    type: "CustomEnqueuedTime",
    consumerGroupName: "cgn",
    eventHubName: "ehn",
    eventSourceResourceId: "somePathInArm",
    keyName: "managementKey",
    kind: "Microsoft.EventHub",
    location: "West US",
    serviceBusNamespace: "sbn",
    sharedAccessKey: "someSecretvalue",
    time: "2017-04-01T19:20:33.2288820Z",
    timestampPropertyName: "someTimestampProperty",
  };
  const credential = new DefaultAzureCredential();
  const client = new TimeSeriesInsightsClient(credential, subscriptionId);
  const result = await client.eventSources.createOrUpdate(
    resourceGroupName,
    environmentName,
    eventSourceName,
    parameters
  );
  console.log(result);
}

eventSourcesCreateEventHubWithCustomEnquedTime().catch(console.error);
