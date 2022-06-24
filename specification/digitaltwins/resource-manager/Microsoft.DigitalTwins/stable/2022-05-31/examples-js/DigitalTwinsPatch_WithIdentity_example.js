const { AzureDigitalTwinsManagementClient } = require("@azure/arm-digitaltwins");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update metadata of DigitalTwinsInstance.
 *
 * @summary Update metadata of DigitalTwinsInstance.
 * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2022-05-31/examples/DigitalTwinsPatch_WithIdentity_example.json
 */
async function patchADigitalTwinsInstanceResourceWithIdentity() {
  const subscriptionId = "50016170-c839-41ba-a724-51e9df440b9e";
  const resourceGroupName = "resRg";
  const resourceName = "myDigitalTwinsService";
  const digitalTwinsPatchDescription = {
    identity: { type: "None" },
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

patchADigitalTwinsInstanceResourceWithIdentity().catch(console.error);
