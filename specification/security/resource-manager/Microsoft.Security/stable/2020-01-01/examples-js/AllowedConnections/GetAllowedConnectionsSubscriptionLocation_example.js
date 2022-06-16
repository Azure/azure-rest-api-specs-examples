const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the list of all possible traffic between resources for the subscription and location.
 *
 * @summary Gets the list of all possible traffic between resources for the subscription and location.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/AllowedConnections/GetAllowedConnectionsSubscriptionLocation_example.json
 */
async function getAllowedConnectionsOnASubscriptionFromSecurityDataLocation() {
  const subscriptionId = "3eeab341-f466-499c-a8be-85427e154bad";
  const ascLocation = "centralus";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.allowedConnections.listByHomeRegion(ascLocation)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllowedConnectionsOnASubscriptionFromSecurityDataLocation().catch(console.error);
