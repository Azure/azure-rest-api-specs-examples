import com.azure.resourcemanager.billing.models.TransferBillingSubscriptionRequestProperties;

/** Samples for BillingSubscriptions ValidateMove. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ValidateSubscriptionMoveSuccess.json
     */
    /**
     * Sample code: SubscriptionMoveValidateSuccess.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void subscriptionMoveValidateSuccess(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingSubscriptions()
            .validateMoveWithResponse(
                "{billingAccountName}",
                new TransferBillingSubscriptionRequestProperties()
                    .withDestinationInvoiceSectionId(
                        "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{newInvoiceSectionName}"),
                com.azure.core.util.Context.NONE);
    }
}
