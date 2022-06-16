import com.azure.core.util.Context;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListFilterByTag.json
     */
    /**
     * Sample code: UsageDetailsListFilterByTag-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void usageDetailsListFilterByTagLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                null,
                "tags eq 'dev:tools'",
                null,
                null,
                null,
                Context.NONE);
    }
}
