const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Content Key Policy in the Media Services account
 *
 * @summary Create or update a Content Key Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/content-key-policies-create-multiple-options.json
 */
async function createsAContentKeyPolicyWithMultipleOptions() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const contentKeyPolicyName = "PolicyCreatedWithMultipleOptions";
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
      {
        name: "widevineoption",
        configuration: {
          odataType: "#Microsoft.Media.ContentKeyPolicyWidevineConfiguration",
          widevineTemplate:
            '{"allowed_track_types":"SD_HD","content_key_specs":[{"track_type":"SD","security_level":1,"required_output_protection":{"hdcp":"HDCP_V2"}}],"policy_overrides":{"can_play":true,"can_persist":true,"can_renew":false}}',
        },
        restriction: {
          odataType: "#Microsoft.Media.ContentKeyPolicyOpenRestriction",
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

createsAContentKeyPolicyWithMultipleOptions().catch(console.error);
