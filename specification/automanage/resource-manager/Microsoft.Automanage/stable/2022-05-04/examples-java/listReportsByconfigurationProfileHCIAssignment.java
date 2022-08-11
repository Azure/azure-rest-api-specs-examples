import com.azure.core.util.Context;

/** Samples for HciReports ListByConfigurationProfileAssignments. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listReportsByconfigurationProfileHCIAssignment.json
     */
    /**
     * Sample code: List reports by HCI configuration profiles assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listReportsByHCIConfigurationProfilesAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .hciReports()
            .listByConfigurationProfileAssignments("myResourceGroupName", "myClusterName", "default", Context.NONE);
    }
}
