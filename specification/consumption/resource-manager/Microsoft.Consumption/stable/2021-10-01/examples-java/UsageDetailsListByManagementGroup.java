import com.azure.core.util.Context;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListByManagementGroup.json
     */
    /**
     * Sample code: ManagementGroupUsageDetailsList-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void managementGroupUsageDetailsListLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list("subscriptions/00000000-0000-0000-0000-000000000000", null, null, null, null, null, Context.NONE);
    }
}
