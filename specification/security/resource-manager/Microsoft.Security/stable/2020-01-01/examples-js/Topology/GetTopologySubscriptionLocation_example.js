const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list that allows to build a topology view of a subscription and location.
 *
 * @summary Gets a list that allows to build a topology view of a subscription and location.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/GetTopologySubscriptionLocation_example.json
 */
async function getTopologyOnASubscriptionFromSecurityDataLocation() {
  const subscriptionId =
    process.env["SECURITY_SUBSCRIPTION_ID"] || "3eeab341-f466-499c-a8be-85427e154bad";
  const ascLocation = "centralus";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.topology.listByHomeRegion(ascLocation)) {
    resArray.push(item);
  }
  console.log(resArray);
}
