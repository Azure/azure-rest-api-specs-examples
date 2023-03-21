const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a DigitalTwinsInstance endpoint.
 *
 * @summary Delete a DigitalTwinsInstance endpoint.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/DigitalTwinsEndpointDelete_WithIdentity_example.json
 */
async function deleteADigitalTwinsInstanceEndpointWithIdentity() {
  const subscriptionId =
    process.env["DIGITALTWINS_SUBSCRIPTION_ID"] || "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = process.env["DIGITALTWINS_RESOURCE_GROUP"] || "resRg";
  const resourceName = "myDigitalTwinsService";
  const endpointName = "myendpoint";
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.digitalTwinsEndpoint.beginDeleteAndWait(
    resourceGroupName,
    resourceName,
    endpointName
  );
  console.log(result);
}
