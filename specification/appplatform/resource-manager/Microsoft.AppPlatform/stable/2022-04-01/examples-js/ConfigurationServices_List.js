const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all resources in a Service.
 *
 * @summary Handles requests to list all resources in a Service.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/ConfigurationServices_List.json
 */
async function configurationServicesList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationServices.list(resourceGroupName, serviceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

configurationServicesList().catch(console.error);
