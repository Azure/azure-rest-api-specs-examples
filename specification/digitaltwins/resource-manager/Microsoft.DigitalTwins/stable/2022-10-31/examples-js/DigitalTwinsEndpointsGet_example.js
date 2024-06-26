const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get DigitalTwinsInstance Endpoints.
 *
 * @summary Get DigitalTwinsInstance Endpoints.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/DigitalTwinsEndpointsGet_example.json
 */
async function getADigitalTwinsInstanceEndpoints() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.digitalTwinsEndpoint.list(resourceGroupName, resourceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getADigitalTwinsInstanceEndpoints().catch(console.error);
