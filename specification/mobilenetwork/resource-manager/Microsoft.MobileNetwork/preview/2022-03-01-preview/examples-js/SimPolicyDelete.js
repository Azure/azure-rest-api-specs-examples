const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified sim policy.
 *
 * @summary Deletes the specified sim policy.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyDelete.json
 */
async function deleteSimPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const mobileNetworkName = "testMobileNetwork";
  const simPolicyName = "testPolicy";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.simPolicies.beginDeleteAndWait(
    resourceGroupName,
    mobileNetworkName,
    simPolicyName
  );
  console.log(result);
}

deleteSimPolicy().catch(console.error);
