
/**
 * Samples for ApplyUpdates CreateOrUpdateParent.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ApplyUpdates_CreateOrUpdateParent.json
     */
    /**
     * Sample code: ApplyUpdates_CreateOrUpdateParent.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        applyUpdatesCreateOrUpdateParent(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.applyUpdates().createOrUpdateParentWithResponse("examplerg", "Microsoft.Compute",
            "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdvm1", com.azure.core.util.Context.NONE);
    }
}
