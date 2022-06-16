import com.azure.core.util.Context;

/** Samples for LotsOperation ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/LotsListByBillingAccount.json
     */
    /**
     * Sample code: LotsListByBillingAccount.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void lotsListByBillingAccount(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.lotsOperations().listByBillingAccount("1234:5678", null, Context.NONE);
    }
}
