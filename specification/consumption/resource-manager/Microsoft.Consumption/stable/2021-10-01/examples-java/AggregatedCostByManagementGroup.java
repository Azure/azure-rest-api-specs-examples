import com.azure.core.util.Context;

/** Samples for AggregatedCost GetByManagementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/AggregatedCostByManagementGroup.json
     */
    /**
     * Sample code: AggregatedCostByManagementGroup.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void aggregatedCostByManagementGroup(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.aggregatedCosts().getByManagementGroupWithResponse("managementGroupForTest", null, Context.NONE);
    }
}
