import com.azure.resourcemanager.billing.models.TransferProductRequestProperties;

/** Samples for Products Move. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/MoveProduct.json
     */
    /**
     * Sample code: MoveProduct.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void moveProduct(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .products()
            .moveWithResponse(
                "{billingAccountName}",
                "{productName}",
                new TransferProductRequestProperties()
                    .withDestinationInvoiceSectionId(
                        "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{newInvoiceSectionName}"),
                com.azure.core.util.Context.NONE);
    }
}
