const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Claim a random claimable virtual machine in the lab. This operation can take a while to complete.
 *
 * @summary Claim a random claimable virtual machine in the lab. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_ClaimAnyVm.json
 */
async function labsClaimAnyVM() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const name = "{labName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.labs.beginClaimAnyVmAndWait(resourceGroupName, name);
  console.log(result);
}

labsClaimAnyVM().catch(console.error);
