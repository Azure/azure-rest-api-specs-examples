import com.azure.core.util.Context;

/** Samples for HcrpReports ListByConfigurationProfileAssignments. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listReportsByconfigurationProfileHCRPAssignment.json
     */
    /**
     * Sample code: List reports by HCRP configuration profiles assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listReportsByHCRPConfigurationProfilesAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .hcrpReports()
            .listByConfigurationProfileAssignments("myResourceGroupName", "myMachineName", "default", Context.NONE);
    }
}
