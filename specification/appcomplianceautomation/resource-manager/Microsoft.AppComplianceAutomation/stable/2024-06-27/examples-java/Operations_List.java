
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        operationsList(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
