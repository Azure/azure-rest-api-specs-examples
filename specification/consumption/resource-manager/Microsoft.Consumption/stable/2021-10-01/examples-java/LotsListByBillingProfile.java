import com.azure.core.util.Context;

/** Samples for LotsOperation ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/LotsListByBillingProfile.json
     */
    /**
     * Sample code: LotsListByBillingProfile.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void lotsListByBillingProfile(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.lotsOperations().listByBillingProfile("1234:5678", "2468", Context.NONE);
    }
}
