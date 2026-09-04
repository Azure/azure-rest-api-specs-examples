
import com.azure.resourcemanager.providerhub.models.ManifestInfoProperties;

/**
 * Samples for Manifests CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Manifests_CreateOrUpdate.json
     */
    /**
     * Sample code: Manifests_CreateOrUpdate.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void manifestsCreateOrUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.manifests().define("prod").withExistingProviderRegistration("Microsoft.Contoso")
            .withProperties(new ManifestInfoProperties().withManifest("<<Core RP manifest>>")).create();
    }
}
