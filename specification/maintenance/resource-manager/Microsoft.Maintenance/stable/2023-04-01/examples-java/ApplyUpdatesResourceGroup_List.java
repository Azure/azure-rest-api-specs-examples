/** Samples for ApplyUpdateForResourceGroup ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ApplyUpdatesResourceGroup_List.json
     */
    /**
     * Sample code: ApplyUpdatesResourceGroup_List.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void applyUpdatesResourceGroupList(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.applyUpdateForResourceGroups().listByResourceGroup("examplerg", com.azure.core.util.Context.NONE);
    }
}
