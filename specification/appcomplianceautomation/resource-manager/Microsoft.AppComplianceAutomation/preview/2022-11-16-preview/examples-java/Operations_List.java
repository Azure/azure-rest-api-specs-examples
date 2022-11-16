import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void operationsList(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.operations().list(Context.NONE);
    }
}
