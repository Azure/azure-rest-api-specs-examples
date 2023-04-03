const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Service or update an exiting Service.
 *
 * @summary Create a new Service or update an exiting Service.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/Services_CreateOrUpdate_Enterprise.json
 */
async function servicesCreateOrUpdateEnterprise() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const resource = {
    location: "eastus",
    properties: {
      marketplaceResource: {
        plan: "tanzu-asc-ent-mtr",
        product: "azure-spring-cloud-vmware-tanzu-2",
        publisher: "vmware-inc",
      },
    },
    sku: { name: "E0", tier: "Enterprise" },
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
