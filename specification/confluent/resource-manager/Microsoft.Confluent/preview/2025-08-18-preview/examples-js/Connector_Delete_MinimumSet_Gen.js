const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to delete confluent connector by name
 *
 * @summary delete confluent connector by name
 * x-ms-original-file: 2025-08-18-preview/Connector_Delete_MinimumSet_Gen.json
 */
async function connectorDeleteMinimumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "DC34558A-05D3-4370-AED8-75E60B381F94";
  const client = new ConfluentManagementClient(credential, subscriptionId);
  await client.connector.delete(
    "rgconfluent",
    "frwocpndztguhgng",
    "duq",
    "chw",
    "suaugvwtvhexoqdrmxknvyiobq",
  );
}
