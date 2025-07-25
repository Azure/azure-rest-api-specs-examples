const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Description for Updates site's Authentication / Authorization settings for apps via the V2 format
 *
 * @summary Description for Updates site's Authentication / Authorization settings for apps via the V2 format
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/UpdateAuthSettingsV2.json
 */
async function updateAuthSettingsV2() {
  const subscriptionId =
    process.env["APPSERVICE_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPSERVICE_RESOURCE_GROUP"] || "testrg123";
  const name = "sitef6141";
  const siteAuthSettingsV2 = {
    globalValidation: {
      excludedPaths: ["/nosecrets/Path"],
      requireAuthentication: true,
      unauthenticatedClientAction: "Return403",
    },
    httpSettings: {
      forwardProxy: {
        convention: "Standard",
        customHostHeaderName: "authHeader",
        customProtoHeaderName: "customProtoHeader",
      },
      requireHttps: true,
      routes: { apiPrefix: "/authv2/" },
    },
    identityProviders: {
      google: {
        enabled: true,
        login: { scopes: ["admin"] },
        registration: {
          clientId: "42d795a9-8abb-4d06-8534-39528af40f8e.apps.googleusercontent.com",
          clientSecretSettingName: "ClientSecret",
        },
        validation: { allowedAudiences: ["https://example.com"] },
      },
    },
    login: {
      allowedExternalRedirectUrls: ["https://someurl.com"],
      cookieExpiration: {
        convention: "IdentityProviderDerived",
        timeToExpiration: "2022:09-01T00:00Z",
      },
      nonce: { validateNonce: true },
      preserveUrlFragmentsForLogins: true,
      routes: { logoutEndpoint: "https://app.com/logout" },
      tokenStore: {
        enabled: true,
        fileSystem: { directory: "/wwwroot/sites/example" },
        tokenRefreshExtensionHours: 96,
      },
    },
    platform: {
      configFilePath: "/auth/config.json",
      enabled: true,
      runtimeVersion: "~1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.updateAuthSettingsV2(
    resourceGroupName,
    name,
    siteAuthSettingsV2,
  );
  console.log(result);
}
