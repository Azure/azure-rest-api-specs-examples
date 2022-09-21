const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all the storages of one Azure Spring Apps resource.
 *
 * @summary List all the storages of one Azure Spring Apps resource.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/Storages_List.json
 */
async function storagesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myService";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.storages.list(resourceGroupName, serviceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

storagesList().catch(console.error);
