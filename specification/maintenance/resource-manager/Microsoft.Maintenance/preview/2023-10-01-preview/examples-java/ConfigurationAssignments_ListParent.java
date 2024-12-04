
/**
 * Samples for ConfigurationAssignments ListParent.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ConfigurationAssignments_ListParent.json
     */
    /**
     * Sample code: ConfigurationAssignments_ListParent.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        configurationAssignmentsListParent(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignments().listParent("examplerg", "Microsoft.Compute", "virtualMachineScaleSets",
            "smdtest1", "virtualMachines", "smdtestvm1", com.azure.core.util.Context.NONE);
    }
}
