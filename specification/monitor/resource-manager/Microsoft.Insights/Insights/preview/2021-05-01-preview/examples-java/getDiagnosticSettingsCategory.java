
/**
 * Samples for DiagnosticSettingsCategory Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-05-01-preview/getDiagnosticSettingsCategory.json
     */
    /**
     * Sample code: Gets the diagnostic setting.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void getsTheDiagnosticSetting(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDiagnosticSettingsCategories().getWithResponse(
            "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6",
            "WorkflowRuntime", com.azure.core.util.Context.NONE);
    }
}
