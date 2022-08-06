const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing Content Key Policy in the Media Services account
 *
 * @summary Updates an existing Content Key Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-update.json
 */
async function updateAContentKeyPolicy() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const contentKeyPolicyName = "PolicyWithClearKeyOptionAndTokenRestriction";
  const parameters = {
    description: "Updated Policy",
    options: [
      {
        name: "ClearKeyOption",
        configuration: {
          odataType: "#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration",
        },
        restriction: {
          odataType: "#Microsoft.Media.ContentKeyPolicyOpenRestriction",
        },
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.contentKeyPolicies.update(
    resourceGroupName,
    accountName,
    contentKeyPolicyName,
    parameters
  );
  console.log(result);
}

updateAContentKeyPolicy().catch(console.error);
