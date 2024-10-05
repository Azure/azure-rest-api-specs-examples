
/**
 * Samples for Charges List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ChargesListByModernBillingAccount.json
     */
    /**
     * Sample code: ChargesListByBillingAccount-Modern.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void
        chargesListByBillingAccountModern(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.charges().listWithResponse("providers/Microsoft.Billing/billingAccounts/1234:56789", "2019-09-01",
            "2019-10-31", null, null, com.azure.core.util.Context.NONE);
    }
}
