
/**
 * Samples for ConfigurationAssignments GetParent.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ConfigurationAssignments_GetParent.json
     */
    /**
     * Sample code: ConfigurationAssignments_GetParent.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        configurationAssignmentsGetParent(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignments().getParentWithResponse("examplerg", "Microsoft.Compute",
            "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdvm1", "workervmPolicy",
            com.azure.core.util.Context.NONE);
    }
}
