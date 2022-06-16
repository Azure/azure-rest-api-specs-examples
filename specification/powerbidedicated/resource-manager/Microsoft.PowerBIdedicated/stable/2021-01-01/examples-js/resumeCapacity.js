const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Resumes operation of the specified Dedicated capacity instance.
 *
 * @summary Resumes operation of the specified Dedicated capacity instance.
 * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/resumeCapacity.json
 */
async function getDetailsOfACapacity() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const dedicatedCapacityName = "azsdktest";
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const result = await client.capacities.beginResumeAndWait(
    resourceGroupName,
    dedicatedCapacityName
  );
  console.log(result);
}

getDetailsOfACapacity().catch(console.error);
