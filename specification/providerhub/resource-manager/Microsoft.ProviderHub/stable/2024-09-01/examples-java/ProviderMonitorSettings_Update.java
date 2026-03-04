
/**
 * Samples for ProviderMonitorSettings Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ProviderMonitorSettings_Update.json
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
