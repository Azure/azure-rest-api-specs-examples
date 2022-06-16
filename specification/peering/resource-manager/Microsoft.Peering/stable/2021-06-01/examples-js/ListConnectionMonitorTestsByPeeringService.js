const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function listAllConnectionMonitorTestsAssociatedWithThePeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.connectionMonitorTests.listByPeeringService(
    resourceGroupName,
    peeringServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllConnectionMonitorTestsAssociatedWithThePeeringService().catch(console.error);
