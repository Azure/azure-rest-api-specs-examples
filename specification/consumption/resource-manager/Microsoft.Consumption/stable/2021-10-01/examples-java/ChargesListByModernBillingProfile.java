import com.azure.core.util.Context;

/** Samples for Charges List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesListByModernBillingProfile.json
     */
    /**
     * Sample code: ChargesListByBillingProfile-Modern.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void chargesListByBillingProfileModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .charges()
            .listWithResponse(
                "providers/Microsoft.Billing/BillingAccounts/1234:56789/billingProfiles/2460",
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
