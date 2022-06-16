const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new App or update an exiting App.
 *
 * @summary Create a new App or update an exiting App.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Apps_CreateOrUpdate.json
 */
async function appsCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const appResource = {
    identity: {
      type: "SystemAssigned",
      principalId: undefined,
      tenantId: undefined,
    },
    location: "eastus",
    properties: {
      addonConfigs: {
        applicationConfigurationService: {
          resourceId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/configurationServices/myacs",
        },
        serviceRegistry: {
          resourceId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/serviceRegistries/myServiceRegistry",
        },
      },
      enableEndToEndTLS: false,
      fqdn: "myapp.mydomain.com",
      httpsOnly: false,
      loadedCertificates: [
        {
          loadTrustStore: false,
          resourceId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert1",
        },
        {
          loadTrustStore: true,
          resourceId:
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.AppPlatform/Spring/myservice/certificates/mycert2",
        },
      ],
      persistentDisk: { mountPath: "/mypersistentdisk", sizeInGB: 2 },
      public: true,
      temporaryDisk: { mountPath: "/mytemporarydisk", sizeInGB: 2 },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.apps.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    appName,
    appResource
  );
  console.log(result);
}

appsCreateOrUpdate().catch(console.error);
