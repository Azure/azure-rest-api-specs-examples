const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a PostRulesResource
 *
 * @summary Create a PostRulesResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/examples/PostRules_CreateOrUpdate_MaximumSet_Gen.json
 */
async function postRulesCreateOrUpdateMaximumSetGen() {
  const subscriptionId =
    process.env["PALOALTONETWORKSNGFW_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const globalRulestackName = "lrs1";
  const priority = "1";
  const resource = {
    description: "description of post rule",
    actionType: "Allow",
    applications: ["app1"],
    auditComment: "example comment",
    category: { feeds: ["feed"], urlCustom: ["https://microsoft.com"] },
    decryptionRuleType: "SSLOutboundInspection",
    destination: {
      cidrs: ["1.0.0.1/10"],
      countries: ["India"],
      feeds: ["feed"],
      fqdnLists: ["FQDN1"],
      prefixLists: ["PL1"],
    },
    enableLogging: "DISABLED",
    etag: "c18e6eef-ba3e-49ee-8a85-2b36c863a9d0",
    inboundInspectionCertificate: "cert1",
    negateDestination: "TRUE",
    negateSource: "TRUE",
    protocolPortList: ["80"],
    provisioningState: "Accepted",
    ruleName: "postRule1",
    ruleState: "DISABLED",
    source: {
      cidrs: ["1.0.0.1/10"],
      countries: ["India"],
      feeds: ["feed"],
      prefixLists: ["PL1"],
    },
    tags: [{ key: "keyName", value: "value" }],
    protocol: "HTTP",
  };
  const credential = new DefaultAzureCredential();
  const client = new PaloAltoNetworksCloudngfw(credential, subscriptionId);
  const result = await client.postRules.beginCreateOrUpdateAndWait(
    globalRulestackName,
    priority,
    resource
  );
  console.log(result);
}
