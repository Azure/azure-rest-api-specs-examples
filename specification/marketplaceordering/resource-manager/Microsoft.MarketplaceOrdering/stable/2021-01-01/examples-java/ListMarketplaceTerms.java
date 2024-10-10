
/**
 * Samples for MarketplaceAgreements List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/marketplaceordering/resource-manager/Microsoft.MarketplaceOrdering/stable/2021-01-01/examples/
     * ListMarketplaceTerms.json
     */
    /**
     * Sample code: ListMarketplaceTerms.
     * 
     * @param manager Entry point to MarketplaceOrderingManager.
     */
    public static void
        listMarketplaceTerms(com.azure.resourcemanager.marketplaceordering.MarketplaceOrderingManager manager) {
        manager.marketplaceAgreements().listWithResponse(com.azure.core.util.Context.NONE);
    }
}
