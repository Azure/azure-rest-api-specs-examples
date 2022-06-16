import com.azure.core.util.Context;

/** Samples for Marketplaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/MarketplacesByDepartment_ListByBillingPeriod.json
     */
    /**
     * Sample code: DepartmentMarketplacesListForBillingPeriod.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void departmentMarketplacesListForBillingPeriod(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.marketplaces().list("providers/Microsoft.Billing/departments/123456", null, null, null, Context.NONE);
    }
}
