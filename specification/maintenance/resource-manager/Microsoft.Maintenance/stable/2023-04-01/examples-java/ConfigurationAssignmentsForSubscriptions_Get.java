/** Samples for ConfigurationAssignmentsForSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignmentsForSubscriptions_Get.json
     */
    /**
     * Sample code: ConfigurationAssignments_GetParent.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsGetParent(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .configurationAssignmentsForSubscriptions()
            .getWithResponse("workervmConfiguration", com.azure.core.util.Context.NONE);
    }
}
