import com.azure.core.util.Context;

/** Samples for MonitoringSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/MonitoringSettings_Get.json
     */
    /**
     * Sample code: MonitoringSettings_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void monitoringSettingsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getMonitoringSettings()
            .getWithResponse("myResourceGroup", "myservice", Context.NONE);
    }
}
