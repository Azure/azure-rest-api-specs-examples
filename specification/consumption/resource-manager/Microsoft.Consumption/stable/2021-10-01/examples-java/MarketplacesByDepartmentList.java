import com.azure.core.util.Context;

/** Samples for Marketplaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/MarketplacesByDepartmentList.json
     */
    /**
     * Sample code: DepartmentMarketplacesList.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void departmentMarketplacesList(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.marketplaces().list("providers/Microsoft.Billing/departments/123456", null, null, null, Context.NONE);
    }
}
