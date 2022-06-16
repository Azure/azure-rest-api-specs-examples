import com.azure.core.util.Context;

/** Samples for BillingAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountsList.json
     */
    /**
     * Sample code: BillingAccountsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().list(null, Context.NONE);
    }
}
