const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a FqdnListGlobalRulestackResource
 *
 * @summary Create a FqdnListGlobalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/FqdnListGlobalRulestack_CreateOrUpdate_MaximumSet_Gen.json
 */
async function fqdnListGlobalRulestackCreateOrUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const globalRulestackName = "praval";
  const name = "armid1";
  const resource = {
    description: "string",
    auditComment: "string",
    etag: "aaaaaaaaaaaaaaaaaa",
    fqdnList: ["string1", "string2"],
    provisioningState: "Accepted",
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.fqdnListGlobalRulestack.beginCreateOrUpdateAndWait(
    globalRulestackName,
    name,
    resource
  );
  console.log(result);
}
