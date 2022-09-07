const { AzureMediaServices } = require("@azure/arm-mediaservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a Content Key Policy in the Media Services account
 *
 * @summary Create or update a Content Key Policy in the Media Services account
 * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/content-key-policies-create-playready-open.json
 */
async function createsAContentKeyPolicyWithPlayReadyOptionAndOpenRestriction() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "contoso";
  const accountName = "contosomedia";
  const contentKeyPolicyName = "PolicyWithPlayReadyOptionAndOpenRestriction";
  const parameters = {
    description: "ArmPolicyDescription",
    options: [
      {
        name: "ArmPolicyOptionName",
        configuration: {
          odataType: "#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration",
          licenses: [
            {
              allowTestDevices: true,
              beginDate: new Date("2017-10-16T18:22:53.46Z"),
              contentKeyLocation: {
                odataType:
                  "#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader",
              },
              contentType: "UltraVioletDownload",
              licenseType: "Persistent",
              playRight: {
                allowPassingVideoContentToUnknownOutput: "NotAllowed",
                digitalVideoOnlyContentRestriction: false,
                imageConstraintForAnalogComponentVideoRestriction: true,
                imageConstraintForAnalogComputerMonitorRestriction: false,
                scmsRestriction: 2,
              },
              securityLevel: "SL150",
            },
          ],
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

createsAContentKeyPolicyWithPlayReadyOptionAndOpenRestriction().catch(console.error);
