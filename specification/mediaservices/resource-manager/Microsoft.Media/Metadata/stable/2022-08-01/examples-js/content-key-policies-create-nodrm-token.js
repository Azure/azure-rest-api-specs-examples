const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Content Key Policy in the Media Services account
 *
 * @summary Create or update a Content Key Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-create-nodrm-token.json
 */
async function createsAContentKeyPolicyWithClearKeyOptionAndTokenRestriction() {
  const subscriptionId =
    process.env["MEDIASERVICES_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MEDIASERVICES_RESOURCE_GROUP"] || "contoso";
  const accountName = "contosomedia";
  const contentKeyPolicyName = "PolicyWithClearKeyOptionAndSwtTokenRestriction";
  const parameters = {
    description: "ArmPolicyDescription",
    options: [
      {
        name: "ClearKeyOption",
        configuration: {
          odataType: "#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration",
        },
        restriction: {
          odataType: "#Microsoft.Media.ContentKeyPolicyTokenRestriction",
          audience: "urn:audience",
          issuer: "urn:issuer",
          primaryVerificationKey: {
            odataType: "#Microsoft.Media.ContentKeyPolicySymmetricTokenKey",
            keyValue: Buffer.from("AAAAAAAAAAAAAAAAAAAAAA=="),
          },
          restrictionTokenType: "Swt",
        },
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureMediaServices(credential, subscriptionId);
  const result = await client.contentKeyPolicies.createOrUpdate(
    resourceGroupName,
    accountName,
    contentKeyPolicyName,
    parameters
  );
  console.log(result);
}
