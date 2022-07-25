const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks that the data connection parameters are valid.
 *
 * @summary Checks that the data connection parameters are valid.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionValidation.json
 */
async function kustoPoolDataConnectionValidation() {
  const subscriptionId = "12345678-1234-1234-1234-123456789098";
  const resourceGroupName = "kustorptest";
  const workspaceName = "kustorptest";
  const kustoPoolName = "kustoclusterrptest4";
  const databaseName = "KustoDatabase8";
  const parameters = {
    dataConnectionName: "DataConnections8",
    properties: { kind: "EventHub" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.kustoPoolDataConnections.beginDataConnectionValidationAndWait(
    resourceGroupName,
    workspaceName,
    kustoPoolName,
    databaseName,
    parameters
  );
  console.log(result);
}

kustoPoolDataConnectionValidation().catch(console.error);
