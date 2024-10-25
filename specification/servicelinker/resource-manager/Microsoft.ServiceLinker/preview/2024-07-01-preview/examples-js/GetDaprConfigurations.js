const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the dapr configuration supported by Service Connector.
 *
 * @summary List the dapr configuration supported by Service Connector.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/GetDaprConfigurations.json
 */
async function getDaprConfigurations() {
  const resourceUri =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.linkers.listDaprConfigurations(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
