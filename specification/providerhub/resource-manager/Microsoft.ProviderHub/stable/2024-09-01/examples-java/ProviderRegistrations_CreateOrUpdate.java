
import com.azure.resourcemanager.providerhub.fluent.models.ProviderRegistrationInner;
import com.azure.resourcemanager.providerhub.models.CrossTenantTokenValidation;
import com.azure.resourcemanager.providerhub.models.ExpeditedRolloutIntent;
import com.azure.resourcemanager.providerhub.models.ProviderRegistrationProperties;
import com.azure.resourcemanager.providerhub.models.Readiness;
import com.azure.resourcemanager.providerhub.models.ResourceProviderCapabilities;
import com.azure.resourcemanager.providerhub.models.ResourceProviderCapabilitiesEffect;
import com.azure.resourcemanager.providerhub.models.ResourceProviderManagementErrorResponseMessageOptions;
import com.azure.resourcemanager.providerhub.models.ResourceProviderManagementExpeditedRolloutMetadata;
import com.azure.resourcemanager.providerhub.models.ResourceProviderManifestPropertiesManagement;
import com.azure.resourcemanager.providerhub.models.ResourceProviderService;
import com.azure.resourcemanager.providerhub.models.ResourceProviderType;
import com.azure.resourcemanager.providerhub.models.ServerFailureResponseMessageType;
import com.azure.resourcemanager.providerhub.models.ServiceStatus;
import com.azure.resourcemanager.providerhub.models.ServiceTreeInfo;
import java.util.Arrays;

/**
 * Samples for ProviderRegistrations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * ProviderRegistrations_CreateOrUpdate.json
     */
    /**
     * Sample code: ProviderRegistrations_CreateOrUpdate.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        providerRegistrationsCreateOrUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerRegistrations()
            .createOrUpdate("Microsoft.Contoso",
                new ProviderRegistrationInner()
                    .withProperties(
                        new ProviderRegistrationProperties()
                            .withServices(Arrays.asList(new ResourceProviderService().withServiceName("tags")
                                .withStatus(ServiceStatus.INACTIVE)))
                            .withServiceName("root").withProviderVersion(
                                "2.0")
                            .withProviderType(ResourceProviderType.INTERNAL)
                            .withManagement(new ResourceProviderManifestPropertiesManagement()
                                .withIncidentRoutingService("Contoso Resource Provider")
                                .withIncidentRoutingTeam("Contoso Triage")
                                .withIncidentContactEmail("helpme@contoso.com")
                                .withServiceTreeInfos(Arrays
                                    .asList(new ServiceTreeInfo().withServiceId("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69")
                                        .withComponentId("d1b7d8ba-05e2-48e6-90d6-d781b99c6e69")
                                        .withReadiness(Readiness.IN_DEVELOPMENT)))
                                .withExpeditedRolloutSubmitters(Arrays.asList("SPARTA-PlatformServiceOperator"))
                                .withErrorResponseMessageOptions(
                                    new ResourceProviderManagementErrorResponseMessageOptions()
                                        .withServerFailureResponseMessageType(
                                            ServerFailureResponseMessageType.OUTAGE_REPORTING))
                                .withExpeditedRolloutMetadata(new ResourceProviderManagementExpeditedRolloutMetadata()
                                    .withEnabled(false).withExpeditedRolloutIntent(ExpeditedRolloutIntent.HOTFIX))
                                .withCanaryManifestOwners(Arrays.asList("SPARTA-PlatformServiceAdmin"))
                                .withPcCode("fakeTokenPlaceholder").withProfitCenterProgramId("1234"))
                            .withCapabilities(Arrays.asList(
                                new ResourceProviderCapabilities().withQuotaId("CSP_2015-05-01")
                                    .withEffect(ResourceProviderCapabilitiesEffect.ALLOW),
                                new ResourceProviderCapabilities().withQuotaId("CSP_MG_2017-12-01")
                                    .withEffect(ResourceProviderCapabilitiesEffect.ALLOW)))
                            .withCrossTenantTokenValidation(CrossTenantTokenValidation.ENSURE_SECURE_VALIDATION)),
                com.azure.core.util.Context.NONE);
    }
}
