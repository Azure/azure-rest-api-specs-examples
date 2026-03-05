
/**
 * Samples for ProviderMonitorSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ProviderMonitorSettings_Delete.json
     */
    /**
     * Sample code: ProviderMonitorSettings_Delete.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerMonitorSettingsDelete(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerMonitorSettings().deleteByResourceGroupWithResponse("default", "ContosoMonitorSetting",
            com.azure.core.util.Context.NONE);
    }
}
