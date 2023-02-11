/** Samples for ScheduledActions Run. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledAction-sendNow-private.json
     */
    /**
     * Sample code: ScheduledActionSendNow.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void scheduledActionSendNow(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager.scheduledActions().runWithResponse("monthlyCostByResource", com.azure.core.util.Context.NONE);
    }
}
