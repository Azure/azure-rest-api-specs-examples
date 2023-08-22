/** Samples for ConfigurationAssignmentsForSubscriptions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignmentsForSubscriptions_Delete.json
     */
    /**
     * Sample code: ConfigurationAssignmentsForSubscriptions_Delete.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsForSubscriptionsDelete(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .configurationAssignmentsForSubscriptions()
            .deleteWithResponse("workervmConfiguration", com.azure.core.util.Context.NONE);
    }
}
