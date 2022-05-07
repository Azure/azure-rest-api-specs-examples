Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-databoxedge_2.0.1/sdk/databoxedge/arm-databoxedge/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the proactive log collection settings of the specified Data Box Edge/Data Box Gateway device.
 *
 * @summary Gets the proactive log collection settings of the specified Data Box Edge/Data Box Gateway device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/GetDiagnosticProactiveLogCollectionSettings.json
 */
async function getDiagnosticProactiveLogCollectionSettings() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = "GroupForEdgeAutomation";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.diagnosticSettings.getDiagnosticProactiveLogCollectionSettings(
    deviceName,
    resourceGroupName
  );
  console.log(result);
}

getDiagnosticProactiveLogCollectionSettings().catch(console.error);
```
