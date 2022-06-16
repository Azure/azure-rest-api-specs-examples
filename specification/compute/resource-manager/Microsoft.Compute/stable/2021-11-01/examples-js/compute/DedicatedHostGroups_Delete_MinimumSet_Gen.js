const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function dedicatedHostGroupsDeleteMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const hostGroupName = "aaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHostGroups.delete(resourceGroupName, hostGroupName);
  console.log(result);
}

dedicatedHostGroupsDeleteMinimumSetGen().catch(console.error);
