import com.azure.resourcemanager.billing.models.TransferProductRequestProperties;

/** Samples for Products ValidateMove. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ValidateProductMoveFailure.json
     */
    /**
     * Sample code: SubscriptionMoveValidateFailure.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void subscriptionMoveValidateFailure(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .products()
            .validateMoveWithResponse(
                "{billingAccountName}",
                "{productName}",
                new TransferProductRequestProperties()
                    .withDestinationInvoiceSectionId(
                        "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{newInvoiceSectionName}"),
                com.azure.core.util.Context.NONE);
    }
}
