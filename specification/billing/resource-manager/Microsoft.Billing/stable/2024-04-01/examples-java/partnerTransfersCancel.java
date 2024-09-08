
/**
 * Samples for PartnerTransfers Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/partnerTransfersCancel.json
     */
    /**
     * Sample code: PartnerTransferCancel.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void partnerTransferCancel(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.partnerTransfers().cancelWithResponse(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "11111111-1111-1111-1111-111111111111", "aabb123", com.azure.core.util.Context.NONE);
    }
}
