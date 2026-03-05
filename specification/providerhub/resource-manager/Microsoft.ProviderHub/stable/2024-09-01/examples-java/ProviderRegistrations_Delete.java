
/**
 * Samples for ProviderRegistrations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ProviderRegistrations_Delete.json
     */
    /**
     * Sample code: ProviderRegistrations_Delete.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerRegistrationsDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerRegistrations().deleteWithResponse("Microsoft.Contoso", com.azure.core.util.Context.NONE);
    }
}
