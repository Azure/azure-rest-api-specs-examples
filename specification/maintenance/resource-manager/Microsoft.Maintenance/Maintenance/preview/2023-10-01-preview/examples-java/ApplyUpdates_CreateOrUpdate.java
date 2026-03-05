
/**
 * Samples for ApplyUpdates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-10-01-preview/ApplyUpdates_CreateOrUpdate.json
     */
    /**
     * Sample code: ApplyUpdates_CreateOrUpdate.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void applyUpdatesCreateOrUpdate(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.applyUpdates().createOrUpdateWithResponse("examplerg", "Microsoft.Compute", "virtualMachineScaleSets",
            "smdtest1", com.azure.core.util.Context.NONE);
    }
}
