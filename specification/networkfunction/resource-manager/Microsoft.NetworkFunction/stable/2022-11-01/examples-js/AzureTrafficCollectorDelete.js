const { AzureTrafficCollectorClient } = require("@azure/arm-networkfunction");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a specified Azure Traffic Collector resource.
 *
 * @summary Deletes a specified Azure Traffic Collector resource.
 * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/AzureTrafficCollectorDelete.json
 */
async function deleteTrafficCollector() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureTrafficCollectorName = "atc";
  const credential = new DefaultAzureCredential();
  const client = new AzureTrafficCollectorClient(credential, subscriptionId);
  const result = await client.azureTrafficCollectors.beginDeleteAndWait(
    resourceGroupName,
    azureTrafficCollectorName
  );
  console.log(result);
}

deleteTrafficCollector().catch(console.error);
