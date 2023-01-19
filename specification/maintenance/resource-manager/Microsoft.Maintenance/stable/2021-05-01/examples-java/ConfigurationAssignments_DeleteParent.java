/** Samples for ConfigurationAssignments DeleteParent. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/ConfigurationAssignments_DeleteParent.json
     */
    /**
     * Sample code: ConfigurationAssignments_DeleteParent.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsDeleteParent(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .configurationAssignments()
            .deleteParentWithResponse(
                "examplerg",
                "Microsoft.Compute",
                "virtualMachineScaleSets",
                "smdtest1",
                "virtualMachines",
                "smdvm1",
                "workervmConfiguration",
                com.azure.core.util.Context.NONE);
    }
}
