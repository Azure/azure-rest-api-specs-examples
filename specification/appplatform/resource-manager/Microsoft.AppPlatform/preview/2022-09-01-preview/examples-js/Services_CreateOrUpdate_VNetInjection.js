const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new Service or update an exiting Service.
 *
 * @summary Create a new Service or update an exiting Service.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2022-09-01-preview/examples/Services_CreateOrUpdate_VNetInjection.json
 */
async function servicesCreateOrUpdateVNetInjection() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const resource = {
    location: "eastus",
    properties: {
      networkProfile: {
        appNetworkResourceGroup: "my-app-network-rg",
        appSubnetId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/apps",
        ingressConfig: { readTimeoutInSeconds: 300 },
        serviceCidr: "10.8.0.0/16,10.244.0.0/16,10.245.0.1/16",
        serviceRuntimeNetworkResourceGroup: "my-service-runtime-network-rg",
        serviceRuntimeSubnetId:
          "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myVirtualNetwork/subnets/serviceRuntime",
      },
      vnetAddons: { logStreamPublicEndpoint: true },
    },
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

servicesCreateOrUpdateVNetInjection().catch(console.error);
