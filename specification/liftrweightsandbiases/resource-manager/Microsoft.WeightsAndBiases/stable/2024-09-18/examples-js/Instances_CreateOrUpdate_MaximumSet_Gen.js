const { WeightsAndBiasesClient } = require("@azure/arm-weightsandbiases");
const { DefaultAzureCredential } = require("@azure/identity");

async function instancesCreateOrUpdateGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "0BCB047F-CB71-4DFE-B0BD-88519F411C2F";
  const client = new WeightsAndBiasesClient(credential, subscriptionId);
  const result = await client.instances.createOrUpdate("rgopenapi", "myinstance", {
    properties: {
      marketplace: {
        subscriptionId: "00000000-0000-0000-0000-000000000000",
        subscriptionStatus: "PendingFulfillmentStart",
        offerDetails: {
          publisherId: "kf",
          offerId: "rfgoevxeke",
          planId: "ufopn",
          planName: "adysakgqlryufffz",
          termUnit: "dgrkojow",
          termId: "kklscqq",
        },
      },
      user: {
        firstName: "kiiehcojcldrlndoid",
        lastName: "zhkvsfqvthwkfkvgxcruyud",
        emailAddress: "user@outlook.com",
        upn: "rmjpgqchpbw",
        phoneNumber: "cogmqmuwfcpstkwbzgkgo",
      },
      partnerProperties: {
        region: "eastus",
        subdomain: "xkecokwnwtkwnkxfgucmzybzzb",
      },
      singleSignOnProperties: {
        type: "Saml",
        state: "Initial",
        enterpriseAppId: "hkxtmpv",
        url: "iqlemoksqdygqyxpp",
        aadDomains: ["mxnw"],
      },
    },
    identity: { type: "None", userAssignedIdentities: {} },
    tags: {},
    location: "pudewmshbcvbt",
  });
  console.log(result);
}
