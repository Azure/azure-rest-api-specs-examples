Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function listsAllAvailableResourceSkUsWithExtendedLocationInformation() {
  const subscriptionId = "{subscription-id}";
  const includeExtendedLocations = "true";
  const options = { includeExtendedLocations: includeExtendedLocations };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceSkus.list(options)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsAllAvailableResourceSkUsWithExtendedLocationInformation().catch(console.error);
```
