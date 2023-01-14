/** Samples for Alerts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2020-06-01/examples/SingleSubscriptionAlert.json
     */
    /**
     * Sample code: SubscriptionAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void subscriptionAlerts(com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .getWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                "22222222-2222-2222-2222-222222222222",
                com.azure.core.util.Context.NONE);
    }
}
