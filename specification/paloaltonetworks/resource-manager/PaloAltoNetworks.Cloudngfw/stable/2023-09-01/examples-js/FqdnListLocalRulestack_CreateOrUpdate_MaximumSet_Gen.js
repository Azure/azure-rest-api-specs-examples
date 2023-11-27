const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a FqdnListLocalRulestackResource
 *
 * @summary Create a FqdnListLocalRulestackResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/FqdnListLocalRulestack_CreateOrUpdate_MaximumSet_Gen.json
 */
async function fqdnListLocalRulestackCreateOrUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
  const resourceGroupName = process.env["PALOALTONETWORKSNGFW_RESOURCE_GROUP"] || "rgopenapi";
  const localRulestackName = "lrs1";
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
  const result = await client.fqdnListLocalRulestack.beginCreateOrUpdateAndWait(
    resourceGroupName,
    localRulestackName,
    name,
    resource
  );
  console.log(result);
}
