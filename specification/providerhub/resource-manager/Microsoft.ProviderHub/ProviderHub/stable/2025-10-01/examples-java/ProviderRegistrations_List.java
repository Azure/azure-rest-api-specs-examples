
/**
 * Samples for ProviderRegistrations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/ProviderRegistrations_List.json
     */
    /**
     * Sample code: ProviderRegistrations_List.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerRegistrationsList(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerRegistrations().list(com.azure.core.util.Context.NONE);
    }
}
