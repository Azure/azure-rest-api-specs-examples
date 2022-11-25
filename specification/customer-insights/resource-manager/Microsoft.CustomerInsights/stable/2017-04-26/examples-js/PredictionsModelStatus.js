const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the model status of prediction.
 *
 * @summary Creates or updates the model status of prediction.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/PredictionsModelStatus.json
 */
async function predictionsModelStatus() {
  const subscriptionId = "c909e979-ef71-4def-a970-bc7c154db8c5";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const predictionName = "sdktest";
  const parameters = { status: "Training" };
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const result = await client.predictions.modelStatus(
    resourceGroupName,
    hubName,
    predictionName,
    parameters
  );
  console.log(result);
}

predictionsModelStatus().catch(console.error);
