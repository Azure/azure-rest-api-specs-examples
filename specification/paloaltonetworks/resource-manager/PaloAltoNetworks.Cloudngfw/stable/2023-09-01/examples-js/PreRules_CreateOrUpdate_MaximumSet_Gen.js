const { PaloAltoNetworksCloudngfw } = require("@azure/arm-paloaltonetworksngfw");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a PreRulesResource
 *
 * @summary Create a PreRulesResource
 * x-ms-original-file: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/PreRules_CreateOrUpdate_MaximumSet_Gen.json
 */
async function preRulesCreateOrUpdateMaximumSetGen() {
  const globalRulestackName = "lrs1";
  const priority = "1";
  const resource = {
    description: "description of pre rule",
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
    ruleName: "preRule1",
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
  const client = new PaloAltoNetworksCloudngfw(credential);
  const result = await client.preRules.beginCreateOrUpdateAndWait(
    globalRulestackName,
    priority,
    resource
  );
  console.log(result);
}
