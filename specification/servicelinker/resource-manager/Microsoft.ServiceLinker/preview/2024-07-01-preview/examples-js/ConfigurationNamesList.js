const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the configuration names generated by Service Connector for all target, client types, auth types.
 *
 * @summary Lists the configuration names generated by Service Connector for all target, client types, auth types.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/ConfigurationNamesList.json
 */
async function getConfigurationNames() {
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.configurationNames.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
