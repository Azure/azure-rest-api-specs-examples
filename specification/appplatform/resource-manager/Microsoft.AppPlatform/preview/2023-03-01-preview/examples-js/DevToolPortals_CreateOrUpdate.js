const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the default Dev Tool Portal or update the existing Dev Tool Portal.
 *
 * @summary Create the default Dev Tool Portal or update the existing Dev Tool Portal.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/preview/2023-03-01-preview/examples/DevToolPortals_CreateOrUpdate.json
 */
async function devToolPortalsCreateOrUpdate() {
  const subscriptionId =
    process.env["APPPLATFORM_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APPPLATFORM_RESOURCE_GROUP"] || "myResourceGroup";
  const serviceName = "myservice";
  const devToolPortalName = "default";
  const devToolPortalResource = {
    properties: {
      features: {
        applicationAccelerator: { state: "Enabled" },
        applicationLiveView: { state: "Enabled" },
      },
      public: true,
      ssoProperties: {
        clientId: "00000000-0000-0000-0000-000000000000",
        clientSecret: "xxxxx",
        metadataUrl:
          "https://login.microsoft.com/00000000-0000-0000-0000-000000000000/v2.0/.well-known/openid-configuration",
        scopes: ["openid"],
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.devToolPortals.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    devToolPortalName,
    devToolPortalResource
  );
  console.log(result);
}
