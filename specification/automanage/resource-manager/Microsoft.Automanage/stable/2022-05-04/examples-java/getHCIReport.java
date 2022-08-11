import com.azure.core.util.Context;

/** Samples for HciReports Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getHCIReport.json
     */
    /**
     * Sample code: Get a report for a HCI configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getAReportForAHCIConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .hciReports()
            .getWithResponse(
                "myResourceGroupName",
                "myClusterName",
                "default",
                "b4e9ee6b-1717-4ff0-a8d2-e6d72c33d5f4",
                Context.NONE);
    }
}
