const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create Cognitive Services Account. Accounts is a resource group wide resource type. It holds the keys for developer to access intelligent APIs. It's also the resource type for billing.
 *
 * @summary create Cognitive Services Account. Accounts is a resource group wide resource type. It holds the keys for developer to access intelligent APIs. It's also the resource type for billing.
 * x-ms-original-file: 2026-11-01/CreateAccountWithTrustedCertificateSources.json
 */
async function createAnAccountWithPrivateCATrust() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.accounts.create("myResourceGroup", "myAccount", {
    kind: "AIServices",
    location: "eastus",
    sku: { name: "S0" },
    identity: { type: "SystemAssigned" },
    properties: {
      allowProjectManagement: true,
      customSubDomainName: "myAccount",
      trustedCertificateSources: [
        {
          type: "AzureKeyVault",
          resourceId:
            "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/myKeyVault",
          certificateBundleReferences: [
            {
              secretName: "corporate-ca-bundle",
              secretVersion: "11111111111111111111111111111111",
            },
            { secretName: "partner-ca-bundle", secretVersion: "22222222222222222222222222222222" },
          ],
        },
      ],
    },
  });
  console.log(result);
}
