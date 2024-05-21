
/**
 * Samples for ConfigurationAssignments List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ConfigurationAssignments_List.json
     */
    /**
     * Sample code: ConfigurationAssignments_List.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsList(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignments().list("examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1",
            com.azure.core.util.Context.NONE);
    }
}
