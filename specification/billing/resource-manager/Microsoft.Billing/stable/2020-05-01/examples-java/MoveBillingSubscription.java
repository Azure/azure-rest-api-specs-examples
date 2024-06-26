import com.azure.resourcemanager.billing.models.TransferBillingSubscriptionRequestProperties;

/** Samples for BillingSubscriptions Move. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/MoveBillingSubscription.json
     */
    /**
     * Sample code: MoveBillingSubscription.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void moveBillingSubscription(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingSubscriptions()
            .move(
                "{billingAccountName}",
                new TransferBillingSubscriptionRequestProperties()
                    .withDestinationInvoiceSectionId(
                        "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{newInvoiceSectionName}"),
                com.azure.core.util.Context.NONE);
    }
}
