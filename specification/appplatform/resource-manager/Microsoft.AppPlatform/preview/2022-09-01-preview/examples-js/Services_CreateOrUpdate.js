const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Service or update an exiting Service.
 *
 * @summary Create a new Service or update an exiting Service.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/Services_CreateOrUpdate.json
 */
async function servicesCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const resource = {
    location: "eastus",
    properties: {},
    sku: { name: "S0", tier: "Standard" },
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.services.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    resource
  );
  console.log(result);
}

servicesCreateOrUpdate().catch(console.error);
