
/**
 * Samples for ProviderMonitorSettings List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * ProviderMonitorSettings_ListBySubscription.json
     */
    /**
     * Sample code: ProviderMonitorSettings_ListBySubscription.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void
        providerMonitorSettingsListBySubscription(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerMonitorSettings().list(com.azure.core.util.Context.NONE);
    }
}
