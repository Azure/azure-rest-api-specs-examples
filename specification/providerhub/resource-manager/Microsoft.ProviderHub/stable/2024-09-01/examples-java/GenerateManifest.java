
/**
 * Samples for ResourceProvider GenerateManifest.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/GenerateManifest.json
     */
    /**
     * Sample code: GenerateManifest.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void generateManifest(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.resourceProviders().generateManifestWithResponse("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
