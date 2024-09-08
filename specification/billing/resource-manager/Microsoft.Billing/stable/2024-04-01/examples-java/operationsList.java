
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/operationsList.json
     */
    /**
     * Sample code: OperationsList.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void operationsList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
