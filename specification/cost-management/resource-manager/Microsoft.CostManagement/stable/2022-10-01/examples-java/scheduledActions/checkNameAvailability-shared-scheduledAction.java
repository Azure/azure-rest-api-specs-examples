import com.azure.resourcemanager.costmanagement.models.CheckNameAvailabilityRequest;

/** Samples for ScheduledActions CheckNameAvailabilityByScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/checkNameAvailability-shared-scheduledAction.json
     */
    /**
     * Sample code: ScheduledActionCheckNameAvailabilityByScope.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void scheduledActionCheckNameAvailabilityByScope(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .scheduledActions()
            .checkNameAvailabilityByScopeWithResponse(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                new CheckNameAvailabilityRequest()
                    .withName("testName")
                    .withType("Microsoft.CostManagement/ScheduledActions"),
                com.azure.core.util.Context.NONE);
    }
}
