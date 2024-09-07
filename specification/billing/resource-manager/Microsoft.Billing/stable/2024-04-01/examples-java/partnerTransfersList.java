
/**
 * Samples for PartnerTransfers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/partnerTransfersList.json
     */
    /**
     * Sample code: PartnerTransfersList.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void partnerTransfersList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.partnerTransfers().list(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "11111111-1111-1111-1111-111111111111", com.azure.core.util.Context.NONE);
    }
}
