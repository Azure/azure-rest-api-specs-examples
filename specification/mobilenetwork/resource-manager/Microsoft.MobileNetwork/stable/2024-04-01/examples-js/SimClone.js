const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Clone SIMs to another SIM Group
 *
 * @summary Clone SIMs to another SIM Group
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/SimClone.json
 */
async function cloneListOfSiMSToTargetSimGroup() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "testResourceGroupName";
  const simGroupName = "testSimGroup";
  const parameters = {
    sims: ["testSim", "testSim2"],
    targetSimGroupId: {
      id: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg2/providers/Microsoft.MobileNetwork/simGroups/testSimGroup1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.sims.beginCloneAndWait(resourceGroupName, simGroupName, parameters);
  console.log(result);
}
