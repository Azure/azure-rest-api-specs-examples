
/**
 * Samples for ScheduledActions ListByScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/
     * scheduledActions/scheduledActions-list-shared.json
     */
    /**
     * Sample code: ScheduledActionsListByScope.
     * 
     * @param manager Entry point to CostManagementManager.
     */
    public static void
        scheduledActionsListByScope(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.scheduledActions().listByScope("subscriptions/00000000-0000-0000-0000-000000000000", null,
            com.azure.core.util.Context.NONE);
    }
}
