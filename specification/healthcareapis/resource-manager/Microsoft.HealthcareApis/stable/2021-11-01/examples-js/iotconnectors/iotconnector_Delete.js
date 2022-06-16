const { HealthcareApisManagementClient } = require("@azure/arm-healthcareapis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an IoT Connector.
 *
 * @summary Deletes an IoT Connector.
 * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_Delete.json
 */
async function deleteAnIoTConnector() {
  const subscriptionId = "subid";
  const resourceGroupName = "testRG";
  const iotConnectorName = "blue";
  const workspaceName = "workspace1";
  const credential = new DefaultAzureCredential();
  const client = new HealthcareApisManagementClient(credential, subscriptionId);
  const result = await client.iotConnectors.beginDeleteAndWait(
    resourceGroupName,
    iotConnectorName,
    workspaceName
  );
  console.log(result);
}

deleteAnIoTConnector().catch(console.error);
