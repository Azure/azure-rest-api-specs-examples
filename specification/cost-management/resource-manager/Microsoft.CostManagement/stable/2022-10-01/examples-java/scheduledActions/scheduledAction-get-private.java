/** Samples for ScheduledActions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledAction-get-private.json
     */
    /**
     * Sample code: PrivateScheduledAction.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void privateScheduledAction(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.scheduledActions().getWithResponse("monthlyCostByResource", com.azure.core.util.Context.NONE);
    }
}
