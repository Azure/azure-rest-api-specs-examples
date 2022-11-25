const { CustomerInsightsManagementClient } = require("@azure/arm-customerinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the connector mappings in the specified connector.
 *
 * @summary Gets all the connector mappings in the specified connector.
 * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/ConnectorMappingsListByConnector.json
 */
async function connectorMappingsListByConnector() {
  const subscriptionId = "subid";
  const resourceGroupName = "TestHubRG";
  const hubName = "sdkTestHub";
  const connectorName = "testConnector8858";
  const credential = new DefaultAzureCredential();
  const client = new CustomerInsightsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.connectorMappings.listByConnector(
    resourceGroupName,
    hubName,
    connectorName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

connectorMappingsListByConnector().catch(console.error);
