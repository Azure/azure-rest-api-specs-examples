
/**
 * Samples for AggregatedCost GetForBillingPeriodByManagementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * AggregatedCostForBillingPeriodByManagementGroup.json
     */
    /**
     * Sample code: AggregatedCostListForBillingPeriodByManagementGroup.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void aggregatedCostListForBillingPeriodByManagementGroup(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.aggregatedCosts().getForBillingPeriodByManagementGroupWithResponse("managementGroupForTest", "201807",
            com.azure.core.util.Context.NONE);
    }
}
