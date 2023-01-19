/** Samples for Updates ListParent. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2021-05-01/examples/Updates_ListParent.json
     */
    /**
     * Sample code: Updates_ListParent.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void updatesListParent(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .updates()
            .listParent(
                "examplerg",
                "Microsoft.Compute",
                "virtualMachineScaleSets",
                "smdtest1",
                "virtualMachines",
                "1",
                com.azure.core.util.Context.NONE);
    }
}
