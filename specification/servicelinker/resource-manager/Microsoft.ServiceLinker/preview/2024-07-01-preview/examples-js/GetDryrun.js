const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get a dryrun job
 *
 * @summary get a dryrun job
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/GetDryrun.json
 */
async function getDryrun() {
  const resourceUri =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
  const dryrunName = "dryrunName";
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.linkers.getDryrun(resourceUri, dryrunName);
  console.log(result);
}
