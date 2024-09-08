
/**
 * Samples for RecipientTransfers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/recipientTransfersGet.json
     */
    /**
     * Sample code: RecipientTransferGet.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void recipientTransferGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.recipientTransfers().getWithResponse("aabb123", com.azure.core.util.Context.NONE);
    }
}
