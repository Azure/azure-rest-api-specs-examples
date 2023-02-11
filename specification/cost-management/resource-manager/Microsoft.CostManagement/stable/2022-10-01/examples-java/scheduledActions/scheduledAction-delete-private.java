/** Samples for ScheduledActions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledAction-delete-private.json
     */
    /**
     * Sample code: PrivateScheduledActionDelete.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void privateScheduledActionDelete(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.scheduledActions().deleteWithResponse("monthlyCostByResource", com.azure.core.util.Context.NONE);
    }
}
