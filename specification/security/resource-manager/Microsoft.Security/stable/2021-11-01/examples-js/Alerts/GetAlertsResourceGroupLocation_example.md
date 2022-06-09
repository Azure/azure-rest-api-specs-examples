```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the alerts that are associated with the resource group that are stored in a specific location
 *
 * @summary List all the alerts that are associated with the resource group that are stored in a specific location
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-11-01/examples/Alerts/GetAlertsResourceGroupLocation_example.json
 */
async function getSecurityAlertsOnAResourceGroupFromASecurityDataLocation() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "westeurope";
  const resourceGroupName = "myRg1";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.alerts.listResourceGroupLevelByRegion(
    ascLocation,
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getSecurityAlertsOnAResourceGroupFromASecurityDataLocation().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
