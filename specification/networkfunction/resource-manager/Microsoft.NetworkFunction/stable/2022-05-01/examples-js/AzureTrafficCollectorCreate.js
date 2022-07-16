const { AzureTrafficCollectorClient } = require("@azure/arm-networkfunction");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Azure Traffic Collector resource
 *
 * @summary Creates or updates a Azure Traffic Collector resource
 * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/AzureTrafficCollectorCreate.json
 */
async function createATrafficCollector() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureTrafficCollectorName = "atc";
  const location = "West US";
  const tags = { key1: "value1" };
  const options = {
    location,
    tags,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureTrafficCollectorClient(credential, subscriptionId);
  const result = await client.azureTrafficCollectors.beginCreateOrUpdateAndWait(
    resourceGroupName,
    azureTrafficCollectorName,
    options
  );
  console.log(result);
}

createATrafficCollector().catch(console.error);
