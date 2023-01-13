/** Samples for Alerts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/SubscriptionAlerts.json
     */
    /**
     * Sample code: SubscriptionAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void subscriptionAlerts(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .listWithResponse("subscriptions/00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
