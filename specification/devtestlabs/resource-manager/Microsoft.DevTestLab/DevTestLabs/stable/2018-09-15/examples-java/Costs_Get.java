
/**
 * Samples for Costs Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Costs_Get.json
     */
    /**
     * Sample code: Costs_Get.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void costsGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.costs().getWithResponse("resourceGroupName", "{labName}", "targetCost", null,
            com.azure.core.util.Context.NONE);
    }
}
