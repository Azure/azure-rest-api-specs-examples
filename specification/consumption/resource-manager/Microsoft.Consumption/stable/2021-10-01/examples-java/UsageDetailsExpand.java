import com.azure.core.util.Context;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsExpand.json
     */
    /**
     * Sample code: UsageDetailsExpand-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void usageDetailsExpandLegacy(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                "meterDetails,additionalInfo",
                "tags eq 'dev:tools'",
                null,
                1,
                null,
                Context.NONE);
    }
}
