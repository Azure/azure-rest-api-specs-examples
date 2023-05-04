const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a CertificateObjectGlobalRulestackResource
 *
 * @summary Create a CertificateObjectGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/CertificateObjectGlobalRulestack_CreateOrUpdate_MinimumSet_Gen.json
 */
async function certificateObjectGlobalRulestackCreateOrUpdateMinimumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const globalRulestackName = "praval";
  const name = "armid1";
  const resource = {
    certificateSelfSigned: "TRUE",
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.certificateObjectGlobalRulestack.beginCreateOrUpdateAndWait(
    globalRulestackName,
    name,
    resource
  );
  console.log(result);
}
