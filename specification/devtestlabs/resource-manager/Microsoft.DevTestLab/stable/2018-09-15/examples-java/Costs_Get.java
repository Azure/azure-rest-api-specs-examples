/** Samples for Costs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_Get.json
     */
    /**
     * Sample code: Costs_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void costsGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .costs()
            .getWithResponse("resourceGroupName", "{labName}", "targetCost", null, com.azure.core.util.Context.NONE);
    }
}
