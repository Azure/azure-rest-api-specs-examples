const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function listCapacityReservationGroupsInSubscription() {
  const subscriptionId = "{subscription-id}";
  const expand = "virtualMachines/$ref";
  const options = { expand: expand };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.capacityReservationGroups.listBySubscription(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listCapacityReservationGroupsInSubscription().catch(console.error);
