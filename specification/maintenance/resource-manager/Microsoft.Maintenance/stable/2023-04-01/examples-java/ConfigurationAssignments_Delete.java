/** Samples for ConfigurationAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignments_Delete.json
     */
    /**
     * Sample code: ConfigurationAssignments_Delete.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsDelete(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .configurationAssignments()
            .deleteWithResponse(
                "examplerg",
                "Microsoft.Compute",
                "virtualMachineScaleSets",
                "smdtest1",
                "workervmConfiguration",
                com.azure.core.util.Context.NONE);
    }
}
