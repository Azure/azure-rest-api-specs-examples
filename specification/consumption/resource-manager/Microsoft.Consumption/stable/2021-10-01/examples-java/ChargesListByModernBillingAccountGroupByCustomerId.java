
/**
 * Samples for Charges List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * ChargesListByModernBillingAccountGroupByCustomerId.json
     */
    /**
     * Sample code: ChargesListByBillingAccountGroupByCustomerId-Modern.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void chargesListByBillingAccountGroupByCustomerIdModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.charges().listWithResponse("providers/Microsoft.Billing/billingAccounts/1234:56789", "2019-09-01",
            "2019-09-30", null, "groupby((properties/customerId))", com.azure.core.util.Context.NONE);
    }
}
