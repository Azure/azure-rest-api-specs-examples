/** Samples for ScheduledActions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledActions-listWithFilter-private.json
     */
    /**
     * Sample code: PrivateScheduledActionsListFilterByViewId.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void privateScheduledActionsListFilterByViewId(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .scheduledActions()
            .list(
                "properties/viewId eq '/providers/Microsoft.CostManagement/views/swaggerExample'",
                com.azure.core.util.Context.NONE);
    }
}
