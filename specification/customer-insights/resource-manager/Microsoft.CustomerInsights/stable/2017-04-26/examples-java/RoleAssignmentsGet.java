/** Samples for RoleAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/RoleAssignmentsGet.json
     */
    /**
     * Sample code: RoleAssignments_Get.
     *
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void roleAssignmentsGet(com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) {
        manager
            .roleAssignments()
            .getWithResponse("TestHubRG", "sdkTestHub", "assignmentName8976", com.azure.core.util.Context.NONE);
    }
}
