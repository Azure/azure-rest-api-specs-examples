import com.azure.core.util.Context;

/** Samples for RoleAssignmentMetrics GetMetricsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2019-08-01-preview/examples/RoleAssignmentMetrics_GetForSubscription.json
     */
    /**
     * Sample code: Get role assignment metrics for subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getRoleAssignmentMetricsForSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignmentMetrics()
            .getMetricsForSubscriptionWithResponse(Context.NONE);
    }
}
