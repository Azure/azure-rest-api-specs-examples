
/**
 * Samples for DiagnosticSettingsOperation Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-05-01-preview/deleteDiagnosticSetting.json
     */
    /**
     * Sample code: Deletes the diagnostic setting.
     * 
     * @param manager Entry point to MonitorManager.
     */
    public static void deletesTheDiagnosticSetting(com.azure.resourcemanager.monitor.MonitorManager manager) {
        manager.serviceClient().getDiagnosticSettingsOperations().deleteWithResponse(
            "subscriptions/1a66ce04-b633-4a0b-b2bc-a912ec8986a6/resourcegroups/viruela1/providers/microsoft.logic/workflows/viruela6",
            "mysetting", com.azure.core.util.Context.NONE);
    }
}
