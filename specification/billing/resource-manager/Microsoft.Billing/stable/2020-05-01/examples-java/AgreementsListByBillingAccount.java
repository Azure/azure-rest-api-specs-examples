/** Samples for Agreements ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/AgreementsListByBillingAccount.json
     */
    /**
     * Sample code: AgreementsListByBillingAccount.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void agreementsListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.agreements().listByBillingAccount("{billingAccountName}", null, com.azure.core.util.Context.NONE);
    }
}
