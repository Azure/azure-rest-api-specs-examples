const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available runtime versions supported by Microsoft.AppPlatform provider.
 *
 * @summary Lists all of the available runtime versions supported by Microsoft.AppPlatform provider.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/RuntimeVersions_ListRuntimeVersions.json
 */
async function runtimeVersionsListRuntimeVersions() {
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential);
  const result = await client.runtimeVersions.listRuntimeVersions();
  console.log(result);
}
