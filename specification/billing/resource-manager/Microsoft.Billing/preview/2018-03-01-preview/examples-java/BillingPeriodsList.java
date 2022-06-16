import com.azure.core.util.Context;

/** Samples for BillingPeriods List. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/BillingPeriodsList.json
     */
    /**
     * Sample code: BillingPeriodsList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingPeriodsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingPeriods().list(null, null, null, Context.NONE);
    }
}
