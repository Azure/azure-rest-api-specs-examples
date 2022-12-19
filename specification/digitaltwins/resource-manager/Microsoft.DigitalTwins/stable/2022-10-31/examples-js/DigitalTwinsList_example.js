const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all the DigitalTwinsInstances in a subscription.
 *
 * @summary Get all the DigitalTwinsInstances in a subscription.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/DigitalTwinsList_example.json
 */
async function getDigitalTwinsInstanceResourcesBySubscription() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.digitalTwins.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

getDigitalTwinsInstanceResourcesBySubscription().catch(console.error);
