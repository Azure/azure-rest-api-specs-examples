const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function exportLogsWhichContainAllThrottledApiRequestsMadeToComputeResourceProviderWithinTheGivenTimePeriod() {
  const subscriptionId = "{subscription-id}";
  const location = "westus";
  const parameters = {
    blobContainerSasUri: "https://somesasuri",
    fromTime: new Date("2018-01-21T01:54:06.862601Z"),
    groupByClientApplicationId: false,
    groupByOperationName: true,
    groupByResourceName: false,
    groupByUserAgent: false,
    toTime: new Date("2018-01-23T01:54:06.862601Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.logAnalytics.beginExportThrottledRequestsAndWait(
    location,
    parameters
  );
  console.log(result);
}

exportLogsWhichContainAllThrottledApiRequestsMadeToComputeResourceProviderWithinTheGivenTimePeriod().catch(
  console.error
);
