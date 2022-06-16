import com.azure.core.util.Context;

/** Samples for Balances GetByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/BalancesByBillingAccount.json
     */
    /**
     * Sample code: Balances.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void balances(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.balances().getByBillingAccountWithResponse("123456", Context.NONE);
    }
}
