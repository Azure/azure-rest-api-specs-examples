import com.azure.core.util.Context;

/** Samples for PriceSheet GetByBillingPeriod. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/PriceSheetForBillingPeriod.json
     */
    /**
     * Sample code: PriceSheetForBillingPeriod.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void priceSheetForBillingPeriod(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.priceSheets().getByBillingPeriodWithResponse("201801", null, null, null, Context.NONE);
    }
}
