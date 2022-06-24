const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get DigitalTwinsInstances resource.
 *
 * @summary Get DigitalTwinsInstances resource.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/DigitalTwinsGet_WithIdentity_example.json
 */
async function getADigitalTwinsInstanceResourceWithIdentity() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.digitalTwins.get(resourceGroupName, resourceName);
  console.log(result);
}

getADigitalTwinsInstanceResourceWithIdentity().catch(console.error);
