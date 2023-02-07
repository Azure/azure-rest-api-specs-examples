/** Samples for ProviderRegistrations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2020-11-20/examples/ProviderRegistrations_List.json
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
