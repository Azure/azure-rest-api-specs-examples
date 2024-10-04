
/**
 * Samples for LotsOperation ListByCustomer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/LotsListByCustomer.
     * json
     */
    /**
     * Sample code: LotsListByCustomer.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void lotsListByCustomer(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.lotsOperations().listByCustomer("1234:5678", "1234:5678", null, com.azure.core.util.Context.NONE);
    }
}
