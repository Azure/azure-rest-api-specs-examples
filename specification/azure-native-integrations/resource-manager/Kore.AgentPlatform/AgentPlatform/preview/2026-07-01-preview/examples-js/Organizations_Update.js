const { AgentPlatformClient } = require("@azure/arm-koreaiartemis");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update mutable properties of an Organization (long-running).
 *
 * @summary update mutable properties of an Organization (long-running).
 * x-ms-original-file: 2026-07-01-preview/Organizations_Update.json
 */
async function organizationsUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AgentPlatformClient(credential, subscriptionId);
  const result = await client.organizations.update("contoso-rg", "contoso-org", {
    properties: {
      partnerProperties: {
        organizationName: "Contoso AI Platform (Updated)",
        billingEmail: "new-billing@contoso.com",
        ownerEmail: "admin@contoso.com",
      },
    },
  });
  console.log(result);
}
