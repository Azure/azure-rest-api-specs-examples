const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a time series database connection.
 *
 * @summary Delete a time series database connection.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/TimeSeriesDatabaseConnectionsDelete_example.json
 */
async function deleteATimeSeriesDatabaseConnectionForADigitalTwinsInstance() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const timeSeriesDatabaseConnectionName = "myConnection";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.timeSeriesDatabaseConnections.beginDeleteAndWait(
    resourceGroupName,
    resourceName,
    timeSeriesDatabaseConnectionName
  );
  console.log(result);
}

deleteATimeSeriesDatabaseConnectionForADigitalTwinsInstance().catch(console.error);
