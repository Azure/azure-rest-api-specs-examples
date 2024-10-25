const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list dryrun jobs
 *
 * @summary list dryrun jobs
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/ListDryrun.json
 */
async function listDryrun() {
  const resourceUri =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.linkers.listDryrun(resourceUri)) {
    resArray.push(item);
  }
  console.log(resArray);
}
