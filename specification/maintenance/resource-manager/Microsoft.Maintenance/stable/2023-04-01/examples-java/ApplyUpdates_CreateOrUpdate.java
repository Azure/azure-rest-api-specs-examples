/** Samples for ApplyUpdates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ApplyUpdates_CreateOrUpdate.json
     */
    /**
     * Sample code: ApplyUpdates_CreateOrUpdate.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void applyUpdatesCreateOrUpdate(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .applyUpdates()
            .createOrUpdateWithResponse(
                "examplerg",
                "Microsoft.Compute",
                "virtualMachineScaleSets",
                "smdtest1",
                com.azure.core.util.Context.NONE);
    }
}
