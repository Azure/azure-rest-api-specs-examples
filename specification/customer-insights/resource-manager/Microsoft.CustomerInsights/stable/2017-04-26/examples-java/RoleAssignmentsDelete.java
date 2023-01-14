/** Samples for RoleAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RoleAssignmentsDelete.json
     */
    /**
     * Sample code: RoleAssignments_Delete.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void roleAssignmentsDelete(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .roleAssignments()
            .deleteWithResponse("TestHubRG", "sdkTestHub", "assignmentName8976", com.azure.core.util.Context.NONE);
    }
}
