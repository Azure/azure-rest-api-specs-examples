const { ProviderHubClient } = require("@azure/arm-providerhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the provider registration.
 *
 * @summary creates or updates the provider registration.
 * x-ms-original-file: 2024-09-01/ProviderRegistrations_CreateOrUpdate.json
 */
async function providerRegistrationsCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
  const client = new ProviderHubClient(credential, subscriptionId);
  const result = await client.providerRegistrations.createOrUpdate("Microsoft.Contoso", {
    properties: {
      capabilities: [
        { effect: "Allow", quotaId: "CSP_2015-05-01" },
        { effect: "Allow", quotaId: "CSP_MG_2017-12-01" },
      ],
      crossTenantTokenValidation: "EnsureSecureValidation",
      management: {
        canaryManifestOwners: ["SPARTA-PlatformServiceAdmin"],
        errorResponseMessageOptions: { serverFailureResponseMessageType: "OutageReporting" },
        expeditedRolloutMetadata: { enabled: false, expeditedRolloutIntent: "Hotfix" },
        expeditedRolloutSubmitters: ["SPARTA-PlatformServiceOperator"],
        incidentContactEmail: "helpme@contoso.com",
        incidentRoutingService: "Contoso Resource Provider",
        incidentRoutingTeam: "Contoso Triage",
        pcCode: "P1234",
        profitCenterProgramId: "1234",
        serviceTreeInfos: [
          {
            componentId: "d1b7d8ba-05e2-48e6-90d6-d781b99c6e69",
            readiness: "InDevelopment",
            serviceId: "d1b7d8ba-05e2-48e6-90d6-d781b99c6e69",
          },
        ],
      },
      providerType: "Internal",
      providerVersion: "2.0",
      serviceName: "root",
      services: [{ serviceName: "tags", status: "Inactive" }],
    },
  });
  console.log(result);
}
