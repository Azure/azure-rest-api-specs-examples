const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function dedicatedHostGroupsListBySubscriptionMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dedicatedHostGroups.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

dedicatedHostGroupsListBySubscriptionMinimumSetGen().catch(console.error);
