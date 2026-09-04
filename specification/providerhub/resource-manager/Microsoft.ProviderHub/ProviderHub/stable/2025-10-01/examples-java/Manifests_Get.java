
/**
 * Samples for Manifests Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/Manifests_Get.json
     */
    /**
     * Sample code: Manifests_Get.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void manifestsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.manifests().getWithResponse("Microsoft.Contoso", "prod", com.azure.core.util.Context.NONE);
    }
}
