import com.azure.core.util.Context;

/** Samples for Charges List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesListByModernCustomer.json
     */
    /**
     * Sample code: ChargesListByCustomer-Modern.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void chargesListByCustomerModern(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .charges()
            .listWithResponse(
                "providers/Microsoft.Billing/BillingAccounts/1234:56789/customers/67890",
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
