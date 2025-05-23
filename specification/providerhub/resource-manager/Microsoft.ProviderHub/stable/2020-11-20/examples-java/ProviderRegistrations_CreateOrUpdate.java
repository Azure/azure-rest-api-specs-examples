
import com.azure.resourcemanager.providerhub.fluent.models.ProviderRegistrationInner;
import com.azure.resourcemanager.providerhub.models.ProviderRegistrationProperties;
import com.azure.resourcemanager.providerhub.models.ResourceProviderCapabilities;
import com.azure.resourcemanager.providerhub.models.ResourceProviderCapabilitiesEffect;
import com.azure.resourcemanager.providerhub.models.ResourceProviderManifestPropertiesManagement;
import com.azure.resourcemanager.providerhub.models.ResourceProviderType;
import java.util.Arrays;

/**
 * Samples for ProviderRegistrations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/
     * ProviderRegistrations_CreateOrUpdate.json
     */
    /**
     * Sample code: ProviderRegistrations_CreateOrUpdate.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        providerRegistrationsCreateOrUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerRegistrations().createOrUpdate("Microsoft.Contoso",
            new ProviderRegistrationInner().withProperties(new ProviderRegistrationProperties()
                .withProviderVersion("2.0").withProviderType(ResourceProviderType.INTERNAL)
                .withManagement(new ResourceProviderManifestPropertiesManagement()
                    .withIncidentRoutingService("Contoso Resource Provider").withIncidentRoutingTeam("Contoso Triage")
                    .withIncidentContactEmail("helpme@contoso.com"))
                .withCapabilities(Arrays.asList(
                    new ResourceProviderCapabilities().withQuotaId("CSP_2015-05-01")
                        .withEffect(ResourceProviderCapabilitiesEffect.ALLOW),
                    new ResourceProviderCapabilities().withQuotaId("CSP_MG_2017-12-01")
                        .withEffect(ResourceProviderCapabilitiesEffect.ALLOW)))),
            com.azure.core.util.Context.NONE);
    }
}
