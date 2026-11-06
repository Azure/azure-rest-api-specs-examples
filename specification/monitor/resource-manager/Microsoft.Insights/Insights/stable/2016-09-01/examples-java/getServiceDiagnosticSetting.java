
/**
 * Samples for ServiceDiagnosticSettingsOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2016-09-01/getServiceDiagnosticSetting.json
     */
    /**
     * Sample code: Get service diagnostic setting.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getServiceDiagnosticSetting(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getServiceDiagnosticSettingsOperations().getWithResponse(
            "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6",
            com.azure.core.util.Context.NONE);
    }
}
