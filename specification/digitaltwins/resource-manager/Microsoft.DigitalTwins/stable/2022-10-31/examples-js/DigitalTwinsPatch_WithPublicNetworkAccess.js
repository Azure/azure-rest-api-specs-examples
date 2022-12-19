const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update metadata of DigitalTwinsInstance.
 *
 * @summary Update metadata of DigitalTwinsInstance.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-10-31/examples/DigitalTwinsPatch_WithPublicNetworkAccess.json
 */
async function patchADigitalTwinsInstanceResourceWithPublicNetworkAccessProperty() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const digitalTwinsPatchDescription = {
    properties: { publicNetworkAccess: "Disabled" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDigitalTwinsManagementClient(credential, subscriptionId);
  const result = await client.digitalTwins.beginUpdateAndWait(
    resourceGroupName,
    resourceName,
    digitalTwinsPatchDescription
  );
  console.log(result);
}

patchADigitalTwinsInstanceResourceWithPublicNetworkAccessProperty().catch(console.error);
