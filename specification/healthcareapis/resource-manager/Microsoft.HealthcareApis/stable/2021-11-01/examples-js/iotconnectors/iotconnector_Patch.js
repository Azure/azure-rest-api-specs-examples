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
