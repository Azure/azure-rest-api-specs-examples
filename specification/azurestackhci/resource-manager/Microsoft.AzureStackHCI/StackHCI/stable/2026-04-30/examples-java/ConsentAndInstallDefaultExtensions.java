
/**
 * Samples for ArcSettings ConsentAndInstallDefaultExtensions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/ConsentAndInstallDefaultExtensions.json
     */
    /**
     * Sample code: Consent And Install Default Extensions.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        consentAndInstallDefaultExtensions(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().consentAndInstallDefaultExtensionsWithResponse("test-rg", "myCluster", "default",
            com.azure.core.util.Context.NONE);
    }
}
