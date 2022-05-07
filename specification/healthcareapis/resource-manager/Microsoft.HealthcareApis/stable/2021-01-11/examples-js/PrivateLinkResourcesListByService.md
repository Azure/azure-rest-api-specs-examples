Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-healthcareapis_2.1.0/sdk/healthcareapis/arm-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

async function privateLinkResourcesListGroupIds() {
  const subscriptionId = "subid";
  const resourceGroupName = "rgname";
  const resourceName = "service1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.privateLinkResources.listByService(resourceGroupName, resourceName);
  console.log(result);
}

privateLinkResourcesListGroupIds().catch(console.error);
```
