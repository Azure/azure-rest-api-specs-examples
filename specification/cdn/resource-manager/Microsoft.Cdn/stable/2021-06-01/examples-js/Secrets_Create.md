Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cdn_7.0.0/sdk/cdn/arm-cdn/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new Secret within the specified profile.
 *
 * @summary Creates a new Secret within the specified profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Secrets_Create.json
 */
async function secretsCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const secretName = "secret1";
  const secret = {
    parameters: {
      type: "CustomerCertificate",
      secretSource: {
        id: "/subscriptions/subid/resourcegroups/RG/providers/Microsoft.KeyVault/vault/kvName/secrets/certificatename",
      },
      secretVersion: "abcdef1234578900abcdef1234567890",
      useLatestVersion: false,
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.secrets.beginCreateAndWait(
    resourceGroupName,
    profileName,
    secretName,
    secret
  );
  console.log(result);
}

secretsCreate().catch(console.error);
```
