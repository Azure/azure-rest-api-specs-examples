Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appcontainers_1.0.0/sdk/appcontainers/arm-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the AuthConfig for a Container App.
 *
 * @summary Create or update the AuthConfig for a Container App.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/AuthConfigs_CreateOrUpdate.json
 */
async function createOrUpdateContainerAppAuthConfig() {
  const subscriptionId = "651f8027-33e8-4ec4-97b4-f6e9f3dc8744";
  const resourceGroupName = "workerapps-rg-xj";
  const containerAppName = "testcanadacentral";
  const authConfigName = "current";
  const authConfigEnvelope = {
    globalValidation: { unauthenticatedClientAction: "AllowAnonymous" },
    identityProviders: {
      facebook: {
        registration: { appId: "123", appSecretSettingName: "facebook-secret" },
      },
    },
    platform: { enabled: true },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerAppsAuthConfigs.createOrUpdate(
    resourceGroupName,
    containerAppName,
    authConfigName,
    authConfigEnvelope
  );
  console.log(result);
}

createOrUpdateContainerAppAuthConfig().catch(console.error);
```
