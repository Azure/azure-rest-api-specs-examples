/** Samples for RoleAssignments ListByHub. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RoleAssignmentsListByHub.json
     */
    /**
     * Sample code: RoleAssignments_ListByHub.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void roleAssignmentsListByHub(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager.roleAssignments().listByHub("TestHubRG", "sdkTestHub", com.azure.core.util.Context.NONE);
    }
}
