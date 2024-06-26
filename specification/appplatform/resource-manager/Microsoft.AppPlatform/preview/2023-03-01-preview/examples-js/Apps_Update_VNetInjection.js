const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Operation to update an exiting App.
 *
 * @summary Operation to update an exiting App.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/Apps_Update_VNetInjection.json
 */
async function appsUpdateVNetInjection() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const appName = "myapp";
  const appResource = {
    identity: {
      type: "SystemAssigned,UserAssigned",
      principalId: undefined,
      tenantId: undefined,
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourceGroups/samplegroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/id1":
          {
            clientId: undefined,
            principalId: undefined,
          },
        "/subscriptions/00000000000000000000000000000000/resourceGroups/samplegroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/id2":
          {
            clientId: undefined,
            principalId: undefined,
          },
      },
    },
    location: "eastus",
    properties: {
      customPersistentDisks: [
        {
          customPersistentDiskProperties: {
            type: "AzureFileVolume",
            mountOptions: [],
            mountPath: "/mypath1/mypath2",
            shareName: "myFileShare",
          },
          storageId: "myASCStorageID",
        },
      ],
      enableEndToEndTLS: false,
      httpsOnly: false,
      persistentDisk: { mountPath: "/mypersistentdisk", sizeInGB: 2 },
      public: true,
      temporaryDisk: { mountPath: "/mytemporarydisk", sizeInGB: 2 },
      vnetAddons: { publicEndpoint: true },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.apps.beginUpdateAndWait(
    resourceGroupName,
    serviceName,
    appName,
    appResource
  );
  console.log(result);
}
