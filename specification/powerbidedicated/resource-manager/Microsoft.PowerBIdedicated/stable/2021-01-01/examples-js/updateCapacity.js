const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the current state of the specified Dedicated capacity.
 *
 * @summary Updates the current state of the specified Dedicated capacity.
 * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/updateCapacity.json
 */
async function updateCapacityParameters() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const dedicatedCapacityName = "azsdktest";
  const capacityUpdateParameters = {
    administration: {
      members: ["azsdktest@microsoft.com", "azsdktest2@microsoft.com"],
    },
    sku: { name: "A1", tier: "PBIE_Azure" },
    tags: { testKey: "testValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const result = await client.capacities.beginUpdateAndWait(
    resourceGroupName,
    dedicatedCapacityName,
    capacityUpdateParameters
  );
  console.log(result);
}

updateCapacityParameters().catch(console.error);
