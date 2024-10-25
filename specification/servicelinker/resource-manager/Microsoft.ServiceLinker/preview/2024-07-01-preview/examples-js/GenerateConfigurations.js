const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate configurations for a Connector.
 *
 * @summary Generate configurations for a Connector.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/GenerateConfigurations.json
 */
async function generateConfiguration() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SERVICELINKER_RESOURCE_GROUP"] || "test-rg";
  const location = "westus";
  const connectorName = "connectorName";
  const parameters = {
    customizedKeys: { aslDocumentDbConnectionString: "MyConnectionstring" },
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.connector.generateConfigurations(
    subscriptionId,
    resourceGroupName,
    location,
    connectorName,
    options,
  );
  console.log(result);
}
