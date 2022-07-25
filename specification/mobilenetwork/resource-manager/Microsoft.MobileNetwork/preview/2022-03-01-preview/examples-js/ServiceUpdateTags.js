const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update service tags.
 *
 * @summary Update service tags.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/ServiceUpdateTags.json
 */
async function updateServiceTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const serviceName = "TestService";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.services.updateTags(
    resourceGroupName,
    mobileNetworkName,
    serviceName,
    parameters
  );
  console.log(result);
}

updateServiceTags().catch(console.error);
