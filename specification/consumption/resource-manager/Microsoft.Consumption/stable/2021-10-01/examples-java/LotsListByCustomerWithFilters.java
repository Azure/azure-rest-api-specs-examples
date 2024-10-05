
/**
 * Samples for LotsOperation ListByCustomer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * LotsListByCustomerWithFilters.json
     */
    /**
     * Sample code: LotsListByCustomerWithFilter.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void lotsListByCustomerWithFilter(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.lotsOperations().listByCustomer("1234:5678", "1234:5678",
            "status eq 'active' AND source eq 'consumptioncommitment'", com.azure.core.util.Context.NONE);
    }
}
