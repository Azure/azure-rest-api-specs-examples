Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-healthcareapis_2.1.0/sdk/healthcareapis/arm-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

async function listAllServicesInResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgname";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.services.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllServicesInResourceGroup().catch(console.error);
```
