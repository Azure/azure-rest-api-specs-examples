
/**
 * Samples for ProviderMonitorSettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * ProviderMonitorSettings_Update.json
     */
    /**
     * Sample code: ProviderMonitorSettings_Update.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerMonitorSettingsUpdate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerMonitorSettings().updateWithResponse("default", "ContosoMonitorSetting",
            com.azure.core.util.Context.NONE);
    }
}
