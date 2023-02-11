import com.azure.resourcemanager.costmanagement.models.ExternalCloudProviderType;

/** Samples for Alerts ListExternal. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/ExternalSubscriptionAlerts.json
     */
    /**
     * Sample code: ExternalSubscriptionAlerts.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void externalSubscriptionAlerts(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .alerts()
            .listExternalWithResponse(
                ExternalCloudProviderType.EXTERNAL_SUBSCRIPTIONS, "100", com.azure.core.util.Context.NONE);
    }
}
