```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list that allows to build a topology view of a subscription and location.
 *
 * @summary Gets a list that allows to build a topology view of a subscription and location.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/GetTopologySubscriptionLocation_example.json
 */
async function getTopologyOnASubscriptionFromSecurityDataLocation() {
  const subscriptionId = "3eeab341-f466-499c-a8be-85427e154bad";
  const ascLocation = "centralus";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.topology.listByHomeRegion(ascLocation)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getTopologyOnASubscriptionFromSecurityDataLocation().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
