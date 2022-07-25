const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update sim policy tags.
 *
 * @summary Update sim policy tags.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyUpdateTags.json
 */
async function updateSimPolicyTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const simPolicyName = "testPolicy";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.simPolicies.updateTags(
    resourceGroupName,
    mobileNetworkName,
    simPolicyName,
    parameters
  );
  console.log(result);
}

updateSimPolicyTags().catch(console.error);
