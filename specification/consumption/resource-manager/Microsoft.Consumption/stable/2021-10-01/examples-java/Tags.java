
/**
 * Samples for Tags Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/Tags.json
     */
    /**
     * Sample code: Tags_Get.
     * 
     * @param manager Entry point to ConsumptionManager.
     */
    public static void tagsGet(com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager.tags().getWithResponse("providers/Microsoft.CostManagement/billingAccounts/1234",
            com.azure.core.util.Context.NONE);
    }
}
