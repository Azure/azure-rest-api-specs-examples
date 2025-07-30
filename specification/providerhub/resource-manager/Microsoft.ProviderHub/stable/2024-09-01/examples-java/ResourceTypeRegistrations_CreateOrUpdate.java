
import com.azure.resourcemanager.providerhub.models.CrossTenantTokenValidation;
import com.azure.resourcemanager.providerhub.models.Notification;
import com.azure.resourcemanager.providerhub.models.NotificationType;
import com.azure.resourcemanager.providerhub.models.OpenApiConfiguration;
import com.azure.resourcemanager.providerhub.models.OpenApiValidation;
import com.azure.resourcemanager.providerhub.models.OptOutHeaderType;
import com.azure.resourcemanager.providerhub.models.Policy;
import com.azure.resourcemanager.providerhub.models.Readiness;
import com.azure.resourcemanager.providerhub.models.Regionality;
import com.azure.resourcemanager.providerhub.models.ResourceAccessPolicy;
import com.azure.resourcemanager.providerhub.models.ResourceConcurrencyControlOption;
import com.azure.resourcemanager.providerhub.models.ResourceTypeEndpoint;
import com.azure.resourcemanager.providerhub.models.ResourceTypeRegistrationProperties;
import com.azure.resourcemanager.providerhub.models.ResourceTypeRegistrationPropertiesManagement;
import com.azure.resourcemanager.providerhub.models.ResourceTypeRegistrationPropertiesRequestHeaderOptions;
import com.azure.resourcemanager.providerhub.models.ResourceTypeRegistrationPropertiesResourceGraphConfiguration;
import com.azure.resourcemanager.providerhub.models.RoutingType;
import com.azure.resourcemanager.providerhub.models.ServiceTreeInfo;
import com.azure.resourcemanager.providerhub.models.SkipNotifications;
import com.azure.resourcemanager.providerhub.models.SwaggerSpecification;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for ResourceTypeRegistrations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * ResourceTypeRegistrations_CreateOrUpdate.json
     */
    /**
     * Sample code: ResourceTypeRegistrations_CreateOrUpdate.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        resourceTypeRegistrationsCreateOrUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.resourceTypeRegistrations().define("employees").withExistingProviderRegistration("Microsoft.Contoso")
            .withProperties(new ResourceTypeRegistrationProperties().withRoutingType(RoutingType.DEFAULT)
                .withCrossTenantTokenValidation(CrossTenantTokenValidation.ENSURE_SECURE_VALIDATION)
                .withRegionality(Regionality.REGIONAL)
                .withEndpoints(
                    Arrays.asList(new ResourceTypeEndpoint().withApiVersions(Arrays.asList("2020-06-01-preview"))
                        .withLocations(Arrays.asList("West US", "East US", "North Europe"))
                        .withRequiredFeatures(Arrays.asList("<feature flag>"))))
                .withSwaggerSpecifications(Arrays.asList(new SwaggerSpecification()
                    .withApiVersions(Arrays.asList("2020-06-01-preview")).withSwaggerSpecFolderUri(
                        "https://github.com/Azure/azure-rest-api-specs/blob/feature/azure/contoso/specification/contoso/resource-manager/Microsoft.SampleRP/")))
                .withRequestHeaderOptions(new ResourceTypeRegistrationPropertiesRequestHeaderOptions()
                    .withOptOutHeaders(OptOutHeaderType.SYSTEM_DATA_CREATED_BY_LAST_MODIFIED_BY))
                .withResourceConcurrencyControlOptions(mapOf("patch",
                    new ResourceConcurrencyControlOption().withPolicy(Policy.SYNCHRONIZE_BEGIN_EXTENSION), "post",
                    new ResourceConcurrencyControlOption().withPolicy(Policy.SYNCHRONIZE_BEGIN_EXTENSION), "put",
                    new ResourceConcurrencyControlOption().withPolicy(Policy.SYNCHRONIZE_BEGIN_EXTENSION)))
                .withResourceGraphConfiguration(new ResourceTypeRegistrationPropertiesResourceGraphConfiguration()
                    .withEnabled(true).withApiVersion("2019-01-01"))
                .withManagement(new ResourceTypeRegistrationPropertiesManagement()
                    .withManifestOwners(Arrays.asList("SPARTA-PlatformServiceAdministrator"))
                    .withAuthorizationOwners(Arrays.asList("RPAAS-PlatformServiceAdministrator"))
                    .withIncidentRoutingService("").withIncidentRoutingTeam("")
                    .withIncidentContactEmail("helpme@contoso.com")
                    .withServiceTreeInfos(
                        Arrays.asList(new ServiceTreeInfo().withServiceId("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69")
                            .withComponentId("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69")
                            .withReadiness(Readiness.IN_DEVELOPMENT)))
                    .withResourceAccessPolicy(ResourceAccessPolicy.NOT_SPECIFIED))
                .withOpenApiConfiguration(new OpenApiConfiguration()
                    .withValidation(new OpenApiValidation().withAllowNoncompliantCollectionResponse(true)))
                .withMetadata(mapOf()).withNotifications(
                    Arrays.asList(new Notification().withNotificationType(NotificationType.SUBSCRIPTION_NOTIFICATION)
                        .withSkipNotifications(SkipNotifications.DISABLED))))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
