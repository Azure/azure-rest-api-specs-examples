const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a CertificateObjectGlobalRulestackResource
 *
 * @summary Create a CertificateObjectGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/CertificateObjectGlobalRulestack_CreateOrUpdate_MaximumSet_Gen.json
 */
async function certificateObjectGlobalRulestackCreateOrUpdateMaximumSetGen() {
  const globalRulestackName = "praval";
  const name = "armid1";
  const resource = {
    description: "description",
    auditComment: "comment",
    certificateSelfSigned: "TRUE",
    certificateSignerResourceId: "",
    etag: "2bf4a339-294d-4c25-b0b2-ef649e9f5c27",
    provisioningState: "Accepted",
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.certificateObjectGlobalRulestack.beginCreateOrUpdateAndWait(
    globalRulestackName,
    name,
    resource
  );
  console.log(result);
}
