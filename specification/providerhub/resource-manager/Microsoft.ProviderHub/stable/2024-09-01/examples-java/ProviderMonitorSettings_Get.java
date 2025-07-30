
/**
 * Samples for ProviderMonitorSettings GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/providerhub/resource-manager/Microsoft.ProviderHub/stable/2024-09-01/examples/
     * ProviderMonitorSettings_Get.json
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
