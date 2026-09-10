const { ProviderHubClient } = require("@azure/arm-providerhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a resource type.
 *
 * @summary creates or updates a resource type.
 * x-ms-original-file: 2025-10-01/ResourceTypeRegistrations_CreateOrUpdate.json
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
        routingType: "Default",
        regionality: "Regional",
        crossTenantTokenValidation: "EnsureSecureValidation",
        endpoints: [
          {
            apiVersions: ["2020-06-01-preview"],
            locations: ["West US", "East US", "North Europe"],
            requiredFeatures: ["<feature flag>"],
          },
        ],
        resourceConcurrencyControlOptions: {
          put: { policy: "SynchronizeBeginExtension" },
          patch: { policy: "SynchronizeBeginExtension" },
          post: { policy: "SynchronizeBeginExtension" },
        },
        swaggerSpecifications: [
          {
            apiVersions: ["2020-06-01-preview"],
            swaggerSpecFolderUri:
              "https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/",
          },
        ],
        resourceGraphConfiguration: { enabled: true, apiVersion: "2019-01-01" },
        management: {
          manifestOwners: ["Contoso-PlatformServiceAdministrator"],
          authorizationOwners: ["RPAAS-PlatformServiceAdministrator"],
          incidentRoutingService: "",
          incidentRoutingTeam: "",
          incidentContactEmail: "helpme@contoso.com",
          resourceAccessPolicy: "NotSpecified",
        },
        metadata: {},
        notifications: [
          { notificationType: "SubscriptionNotification", skipNotifications: "Disabled" },
        ],
        openApiConfiguration: { validation: { allowNoncompliantCollectionResponse: true } },
        requestHeaderOptions: { optOutHeaders: "SystemDataCreatedByLastModifiedBy" },
        throttlingRules: [
          {
            action: "Microsoft.Foo/checkNameAvailability/write",
            metrics: [{ type: "NumberOfRequests", bucketSize: "XLarge", limit: 1 }],
          },
        ],
        privateEndpointConfiguration: {
          minApiVersion: "2022-10-01",
          groupConnectivityInformation: [
            {
              groupId: "Sql",
              requiredMembers: ["Sql_Member"],
              requiredZoneNames: ["Zone"],
              redirectMapId: "test",
            },
          ],
        },
        writeLock: { state: "Enabled" },
        marketplaceType: "ProviderHub",
        resourceManagementOptions: {
          batchProvisioningSupport: {
            maxBatchSize: 10,
            actionConfigurations: [
              { authorizationAction: "Microsoft.Contoso/authorize", maxBatchSize: 5 },
            ],
            batchContractVersion: "2020-06-01-preview",
            maxNestedBatchSize: 5,
            requiredFeatures: ["Microsoft.Contoso/feature1"],
            supportedOperations: "Get",
          },
        },
      },
    },
  );
  console.log(result);
}
