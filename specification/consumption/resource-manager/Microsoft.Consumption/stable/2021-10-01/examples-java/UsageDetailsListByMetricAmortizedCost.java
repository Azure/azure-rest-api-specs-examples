import com.azure.core.util.Context;
import com.azure.resourcemanager.consumption.models.Metrictype;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListByMetricAmortizedCost.json
     */
    /**
     * Sample code: UsageDetailsListByMetricAmortizedCost-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void usageDetailsListByMetricAmortizedCostLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list(
                "subscriptions/00000000-0000-0000-0000-000000000000",
                null,
                null,
                null,
                null,
                Metrictype.AMORTIZEDCOST,
                Context.NONE);
    }
}
