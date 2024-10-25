const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a Linker.
 *
 * @summary Delete a Linker.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/DeleteLinker.json
 */
async function deleteLinker() {
  const resourceUri =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
  const linkerName = "linkName";
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.linker.beginDeleteAndWait(resourceUri, linkerName);
  console.log(result);
}
