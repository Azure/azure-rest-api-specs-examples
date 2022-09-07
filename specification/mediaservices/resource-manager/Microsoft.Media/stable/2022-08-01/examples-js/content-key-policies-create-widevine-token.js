const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Content Key Policy in the Media Services account
 *
 * @summary Create or update a Content Key Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/content-key-policies-create-widevine-token.json
 */
async function createsAContentKeyPolicyWithWidevineOptionAndTokenRestriction() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const contentKeyPolicyName = "PolicyWithWidevineOptionAndJwtTokenRestriction";
  const parameters = {
    description: "ArmPolicyDescription",
    options: [
      {
        name: "widevineoption",
        configuration: {
          odataType: "#Microsoft.Media.ContentKeyPolicyWidevineConfiguration",
          widevineTemplate:
            '{"allowed_track_types":"SD_HD","content_key_specs":[{"track_type":"SD","security_level":1,"required_output_protection":{"hdcp":"HDCP_V2"}}],"policy_overrides":{"can_play":true,"can_persist":true,"can_renew":false}}',
        },
        restriction: {
          odataType: "#Microsoft.Media.ContentKeyPolicyTokenRestriction",
          alternateVerificationKeys: [
            {
              odataType: "#Microsoft.Media.ContentKeyPolicySymmetricTokenKey",
              keyValue: Buffer.from("AAAAAAAAAAAAAAAAAAAAAA=="),
            },
          ],
          audience: "urn:audience",
          issuer: "urn:issuer",
          primaryVerificationKey: {
            odataType: "#Microsoft.Media.ContentKeyPolicyRsaTokenKey",
            exponent: Buffer.from("AQAB"),
            modulus: Buffer.from("AQAD"),
          },
          restrictionTokenType: "Jwt",
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

createsAContentKeyPolicyWithWidevineOptionAndTokenRestriction().catch(console.error);
