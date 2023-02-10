/** Samples for ScheduledActions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledActions-list-private.json
     */
    /**
     * Sample code: PrivateScheduledActionsList.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void privateScheduledActionsList(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.scheduledActions().list(null, com.azure.core.util.Context.NONE);
    }
}
