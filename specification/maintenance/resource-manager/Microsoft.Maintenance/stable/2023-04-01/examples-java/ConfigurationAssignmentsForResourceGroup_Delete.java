/** Samples for ConfigurationAssignmentsForResourceGroup Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignmentsForResourceGroup_Delete.json
     */
    /**
     * Sample code: ConfigurationAssignmentsForResourceGroup_Delete.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsForResourceGroupDelete(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .configurationAssignmentsForResourceGroups()
            .deleteByResourceGroupWithResponse("examplerg", "workervmConfiguration", com.azure.core.util.Context.NONE);
    }
}
