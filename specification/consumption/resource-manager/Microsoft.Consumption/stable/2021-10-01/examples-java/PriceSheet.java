import com.azure.core.util.Context;

/** Samples for PriceSheet Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/PriceSheet.json
     */
    /**
     * Sample code: PriceSheet.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void priceSheet(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.priceSheets().getWithResponse(null, null, null, Context.NONE);
    }
}
