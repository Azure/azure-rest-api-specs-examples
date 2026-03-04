
/**
 * Samples for ConfigurationAssignments DeleteParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-10-01-preview/ConfigurationAssignments_DeleteParent.json
     */
    /**
     * Sample code: ConfigurationAssignments_DeleteParent.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        configurationAssignmentsDeleteParent(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignments().deleteParentWithResponse("examplerg", "Microsoft.Compute",
            "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdvm1", "workervmConfiguration",
            com.azure.core.util.Context.NONE);
    }
}
