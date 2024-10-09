
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/
     * OperationsList.json
     */
    /**
     * Sample code: List Operations.
     * 
     * @param manager Entry point to MarketplaceOrderingManager.
     */
    public static void
        listOperations(com.azure.resourcemanager.marketplaceordering.MarketplaceOrderingManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
