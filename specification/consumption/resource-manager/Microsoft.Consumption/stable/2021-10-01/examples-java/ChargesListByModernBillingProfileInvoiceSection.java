import com.azure.core.util.Context;

/** Samples for Charges List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/ChargesListByModernBillingProfileInvoiceSection.json
     */
    /**
     * Sample code: ChargesListByBillingProfileInvoiceSection-Modern.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void chargesListByBillingProfileInvoiceSectionModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .charges()
            .listWithResponse(
                "providers/Microsoft.Billing/billingAccounts/1234:56789/billingProfiles/42425/invoiceSections/67890",
                "2019-09-01",
                "2019-10-31",
                null,
                null,
                Context.NONE);
    }
}
