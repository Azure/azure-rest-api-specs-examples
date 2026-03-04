
/**
 * Samples for ProviderMonitorSettings GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ProviderMonitorSettings_Get.json
     */
    /**
     * Sample code: ProviderMonitorSettings_Get.
     * 
     * @param manager Entry point to ProviderHubManager.
     */
    public static void providerMonitorSettingsGet(com.azure.resourcemanager.providerhub.ProviderHubManager manager) {
        manager.providerMonitorSettings().getByResourceGroupWithResponse("default", "ContosoMonitorSetting",
            com.azure.core.util.Context.NONE);
    }
}
