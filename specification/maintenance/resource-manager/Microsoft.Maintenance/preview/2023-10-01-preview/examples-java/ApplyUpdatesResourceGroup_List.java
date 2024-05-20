
/**
 * Samples for ApplyUpdateForResourceGroup ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ApplyUpdatesResourceGroup_List.json
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
