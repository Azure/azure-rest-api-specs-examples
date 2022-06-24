const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all existing time series database connections for this DigitalTwins instance.
 *
 * @summary Get all existing time series database connections for this DigitalTwins instance.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/TimeSeriesDatabaseConnectionsList_example.json
 */
async function listTimeSeriesDatabaseConnectionsForADigitalTwinsInstance() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.timeSeriesDatabaseConnections.list(
    resourceGroupName,
    resourceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listTimeSeriesDatabaseConnectionsForADigitalTwinsInstance().catch(console.error);
