const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

async function checkNameAvailabilityOfACapacity() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const location = "West US";
  const capacityParameters = {
    name: "azsdktest",
    type: "Microsoft.PowerBIDedicated/capacities",
  };
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const result = await client.capacities.checkNameAvailability(location, capacityParameters);
  console.log(result);
}

checkNameAvailabilityOfACapacity().catch(console.error);
