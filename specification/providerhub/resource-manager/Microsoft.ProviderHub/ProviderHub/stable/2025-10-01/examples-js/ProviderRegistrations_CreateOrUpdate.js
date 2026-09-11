const { ProviderHubClient } = require("@azure/arm-providerhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the provider registration.
 *
 * @summary creates or updates the provider registration.
 * x-ms-original-file: 2025-10-01/ProviderRegistrations_CreateOrUpdate.json
 */
async function providerRegistrationsCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
  const client = new ProviderHubClient(credential, subscriptionId);
  const result = await client.providerRegistrations.createOrUpdate("Microsoft.Contoso", {
    properties: {
      providerType: "Internal",
      providerVersion: "2.0",
      serviceName: "root",
      services: [{ serviceName: "tags", status: "Inactive" }],
      crossTenantTokenValidation: "EnsureSecureValidation",
      management: {
        incidentRoutingService: "Contoso Resource Provider",
        incidentRoutingTeam: "Contoso Triage",
        incidentContactEmail: "helpme@contoso.com",
        expeditedRolloutSubmitters: ["Contoso-PlatformServiceOperator"],
        expeditedRolloutMetadata: { enabled: false, expeditedRolloutIntent: "Hotfix" },
        errorResponseMessageOptions: { serverFailureResponseMessageType: "OutageReporting" },
        canaryManifestOwners: ["Contoso-PlatformServiceAdmin"],
        pcCode: "P1234",
        profitCenterProgramId: "1234",
      },
      capabilities: [
        { quotaId: "CSP_2015-05-01", effect: "Allow" },
        { quotaId: "CSP_MG_2017-12-01", effect: "Allow" },
      ],
    },
  });
  console.log(result);
}
