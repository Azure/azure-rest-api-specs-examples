
/**
 * Samples for RecipientTransfers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/recipientTransfersList.json
     */
    /**
     * Sample code: RecipientTransfersList.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void recipientTransfersList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.recipientTransfers().list(com.azure.core.util.Context.NONE);
    }
}
