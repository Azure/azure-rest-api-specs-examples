import com.azure.core.util.Context;

/** Samples for BillingProfiles Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileWithExpand.json
     */
    /**
     * Sample code: BillingProfileWithExpand.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfileWithExpand(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingProfiles()
            .getWithResponse("{billingAccountName}", "{billingProfileName}", "invoiceSections", Context.NONE);
    }
}
