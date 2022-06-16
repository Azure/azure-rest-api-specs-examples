import com.azure.core.util.Context;
import com.azure.resourcemanager.billing.models.TransferProductRequestProperties;

/** Samples for Products ValidateMove. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ValidateProductMoveSuccess.json
     */
    /**
     * Sample code: SubscriptionMoveValidateSuccess.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void subscriptionMoveValidateSuccess(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .products()
            .validateMoveWithResponse(
                "{billingAccountName}",
                "{productName}",
                new TransferProductRequestProperties()
                    .withDestinationInvoiceSectionId(
                        "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{newInvoiceSectionName}"),
                Context.NONE);
    }
}
