
/**
 * Samples for Updates List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/Updates_List
     * .json
     */
    /**
     * Sample code: Updates_List.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void updatesList(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.updates().list("examplerg", "Microsoft.Compute", "virtualMachineScaleSets", "smdtest1",
            com.azure.core.util.Context.NONE);
    }
}
