Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-healthcareapis_2.1.0/sdk/healthcareapis/arm-healthcareapis/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patch an IoT Connector.
 *
 * @summary Patch an IoT Connector.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_Patch.json
 */
async function patchAnIoTConnector() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const iotConnectorName = "blue";
  const workspaceName = "workspace1";
  const iotConnectorPatchResource = {
    identity: { type: "SystemAssigned" },
    tags: {
      additionalProp1: "string",
      additionalProp2: "string",
      additionalProp3: "string",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.iotConnectors.beginUpdateAndWait(
    resourceGroupName,
    iotConnectorName,
    workspaceName,
    iotConnectorPatchResource
  );
  console.log(result);
}

patchAnIoTConnector().catch(console.error);
```
