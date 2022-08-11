import com.azure.core.util.Context;

/** Samples for Reports ListByConfigurationProfileAssignments. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listReportsByconfigurationProfileAssignment.json
     */
    /**
     * Sample code: List reports by configuration profiles assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listReportsByConfigurationProfilesAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .reports()
            .listByConfigurationProfileAssignments("myResourceGroupName", "default", "myVMName", Context.NONE);
    }
}
