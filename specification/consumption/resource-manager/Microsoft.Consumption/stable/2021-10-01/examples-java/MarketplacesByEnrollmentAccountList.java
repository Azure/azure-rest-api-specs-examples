
/**
 * Samples for Marketplaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * MarketplacesByEnrollmentAccountList.json
     */
    /**
     * Sample code: EnrollmentAccountMarketplacesList.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void
        enrollmentAccountMarketplacesList(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.marketplaces().list("providers/Microsoft.Billing/enrollmentAccounts/123456", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
