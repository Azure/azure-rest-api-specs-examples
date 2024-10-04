
/**
 * Samples for Marketplaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * MarketplacesByManagementGroup_ListForBillingPeriod.json
     */
    /**
     * Sample code: ManagementGroupMarketplacesListForBillingPeriod.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void managementGroupMarketplacesListForBillingPeriod(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.marketplaces().list("subscriptions/00000000-0000-0000-0000-000000000000", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
