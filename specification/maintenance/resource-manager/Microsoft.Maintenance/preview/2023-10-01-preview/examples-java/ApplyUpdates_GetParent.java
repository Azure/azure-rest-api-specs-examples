
/**
 * Samples for ApplyUpdates GetParent.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ApplyUpdates_GetParent.json
     */
    /**
     * Sample code: ApplyUpdates_GetParent.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void applyUpdatesGetParent(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.applyUpdates().getParentWithResponse("examplerg", "Microsoft.Compute", "virtualMachineScaleSets",
            "smdtest1", "virtualMachines", "smdvm1", "e9b9685d-78e4-44c4-a81c-64a14f9b87b6",
            com.azure.core.util.Context.NONE);
    }
}
