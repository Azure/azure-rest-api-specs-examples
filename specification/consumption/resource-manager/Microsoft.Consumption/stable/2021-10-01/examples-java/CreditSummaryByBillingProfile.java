
/**
 * Samples for Credits Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/
     * CreditSummaryByBillingProfile.json
     */
    /**
     * Sample code: CreditSummaryByBillingProfile.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void creditSummaryByBillingProfile(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.credits().getWithResponse("1234:5678", "2468", com.azure.core.util.Context.NONE);
    }
}
