const { ServiceLinkerManagementClient } = require("@azure/arm-servicelinker");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generate configurations for a Linker.
 *
 * @summary Generate configurations for a Linker.
 * x-ms-original-file: specification/servicelinker/resource-manager/Microsoft.ServiceLinker/preview/2024-07-01-preview/examples/LinkerGenerateConfigurations.json
 */
async function generateConfiguration() {
  const resourceUri =
    "subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Web/sites/test-app";
  const linkerName = "linkName";
  const parameters = {
    customizedKeys: { aslDocumentDbConnectionString: "MyConnectionstring" },
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ServiceLinkerManagementClient(credential);
  const result = await client.linkers.generateConfigurations(resourceUri, linkerName, options);
  console.log(result);
}
