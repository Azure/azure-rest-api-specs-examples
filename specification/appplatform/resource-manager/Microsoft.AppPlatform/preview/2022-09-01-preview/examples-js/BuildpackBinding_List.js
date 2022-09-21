const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Handles requests to list all buildpack bindings in a builder.
 *
 * @summary Handles requests to list all buildpack bindings in a builder.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/BuildpackBinding_List.json
 */
async function buildpackBindingGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const buildServiceName = "default";
  const builderName = "default";
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.buildpackBinding.list(
    resourceGroupName,
    serviceName,
    buildServiceName,
    builderName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

buildpackBindingGet().catch(console.error);
