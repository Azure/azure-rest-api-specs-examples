import com.azure.core.util.Context;

/** Samples for LotsOperation ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/LotsListByBillingAccountWithFilters.json
     */
    /**
     * Sample code: LotsListByBillingAccountWithStatusFilter.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void lotsListByBillingAccountWithStatusFilter(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .lotsOperations()
            .listByBillingAccount(
                "1234:5678", "status eq 'active' AND source eq 'consumptioncommitment'", Context.NONE);
    }
}
