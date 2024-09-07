
import com.azure.resourcemanager.billing.models.InitiateTransferRequest;

/**
 * Samples for Transfers Initiate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/transfersInitiate.json
     */
    /**
     * Sample code: InitiateTransfer.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void initiateTransfer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.transfers().initiateWithResponse(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "yyyy-yyyy-yyy-yyy", "aabb123", new InitiateTransferRequest().withRecipientEmailId("user@contoso.com"),
            com.azure.core.util.Context.NONE);
    }
}
