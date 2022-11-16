import com.azure.core.util.Context;

/** Samples for Reports List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Reports_List.json
     */
    /**
     * Sample code: Reports_List.
     *
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void reportsList(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager
            .reports()
            .list(
                "1",
                100,
                null,
                "00000000-0000-0000-0000-000000000000",
                "00000000-0000-0000-0000-000000000000",
                Context.NONE);
    }
}
