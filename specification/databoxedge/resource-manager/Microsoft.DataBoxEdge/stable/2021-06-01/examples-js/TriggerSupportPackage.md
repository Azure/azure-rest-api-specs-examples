Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-databoxedge_2.0.1/sdk/databoxedge/arm-databoxedge/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Triggers support package on the device
 *
 * @summary Triggers support package on the device
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/TriggerSupportPackage.json
 */
async function triggerSupportPackage() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = "GroupForEdgeAutomation";
  const triggerSupportPackageRequest = {
    include: "DefaultWithDumps",
    maximumTimeStamp: new Date("2018-12-18T02:19:51.4270267Z"),
    minimumTimeStamp: new Date("2018-12-18T02:18:51.4270267Z"),
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.supportPackages.beginTriggerSupportPackageAndWait(
    deviceName,
    resourceGroupName,
    triggerSupportPackageRequest
  );
  console.log(result);
}

triggerSupportPackage().catch(console.error);
```
