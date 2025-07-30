
/**
 * Samples for ProviderMonitorSettings Create.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * ProviderMonitorSettings_Create.json
     */
    /**
     * Sample code: ProviderMonitorSettings_Create.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerMonitorSettingsCreate(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerMonitorSettings().define("ContosoMonitorSetting").withRegion("eastus")
            .withExistingResourceGroup("default").create();
    }
}
