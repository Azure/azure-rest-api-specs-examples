/** Samples for AvailableBalances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AvailableBalanceByBillingProfile.json
     */
    /**
     * Sample code: AvailableBalanceByBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void availableBalanceByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .availableBalances()
            .getWithResponse("{billingAccountName}", "{billingProfileName}", com.azure.core.util.Context.NONE);
    }
}
