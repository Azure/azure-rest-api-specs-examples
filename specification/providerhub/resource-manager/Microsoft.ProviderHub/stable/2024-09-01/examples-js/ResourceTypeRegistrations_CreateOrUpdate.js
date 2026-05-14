const { ProviderHubClient } = require("@azure/arm-providerhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a resource type.
 *
 * @summary creates or updates a resource type.
 * x-ms-original-file: 2024-09-01/ResourceTypeRegistrations_CreateOrUpdate.json
 */
async function resourceTypeRegistrationsCreateOrUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "ab7a8701-f7ef-471a-a2f4-d0ebbf494f77";
  const client = new ProviderHubClient(credential, subscriptionId);
  const result = await client.resourceTypeRegistrations.createOrUpdate(
    "Microsoft.Contoso",
    "employees",
    {
      properties: {
        crossTenantTokenValidation: "EnsureSecureValidation",
        endpoints: [
          {
            apiVersions: ["2020-06-01-preview"],
            locations: ["West US", "East US", "North Europe"],
            requiredFeatures: ["<feature flag>"],
          },
        ],
        management: {
          authorizationOwners: ["RPAAS-PlatformServiceAdministrator"],
          incidentContactEmail: "helpme@contoso.com",
          incidentRoutingService: "",
          incidentRoutingTeam: "",
          manifestOwners: ["SPARTA-PlatformServiceAdministrator"],
          resourceAccessPolicy: "NotSpecified",
          serviceTreeInfos: [
            {
              componentId: "d1b7d8ba-05e2-48e6-90d6-d781b99c6e69",
              readiness: "InDevelopment",
              serviceId: "d1b7d8ba-05e2-48e6-90d6-d781b99c6e69",
            },
          ],
        },
        metadata: {},
        notifications: [
          { notificationType: "SubscriptionNotification", skipNotifications: "Disabled" },
        ],
        openApiConfiguration: { validation: { allowNoncompliantCollectionResponse: true } },
        regionality: "Regional",
        requestHeaderOptions: { optOutHeaders: "SystemDataCreatedByLastModifiedBy" },
        resourceConcurrencyControlOptions: {
          patch: { policy: "SynchronizeBeginExtension" },
          post: { policy: "SynchronizeBeginExtension" },
          put: { policy: "SynchronizeBeginExtension" },
        },
        resourceGraphConfiguration: { apiVersion: "2019-01-01", enabled: true },
        routingType: "Default",
        swaggerSpecifications: [
          {
            apiVersions: ["2020-06-01-preview"],
            swaggerSpecFolderUri:
              "https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/",
          },
        ],
      },
    },
  );
  console.log(result);
}
