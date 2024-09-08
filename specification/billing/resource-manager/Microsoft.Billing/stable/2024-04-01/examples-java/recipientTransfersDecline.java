
/**
 * Samples for RecipientTransfers Decline.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/recipientTransfersDecline.
     * json
     */
    /**
     * Sample code: DeclineTransfer.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void declineTransfer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.recipientTransfers().declineWithResponse("aabb123", com.azure.core.util.Context.NONE);
    }
}
