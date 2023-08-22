/** Samples for ConfigurationAssignmentsWithinSubscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignmentsResultWithinSubscription_List.json
     */
    /**
     * Sample code: ConfigurationAssignmentsResultWithinSubscription_List.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsResultWithinSubscriptionList(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignmentsWithinSubscriptions().list(com.azure.core.util.Context.NONE);
    }
}
