/** Samples for Labs GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_Get.json
     */
    /**
     * Sample code: Labs_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .labs()
            .getByResourceGroupWithResponse("resourceGroupName", "{labName}", null, com.azure.core.util.Context.NONE);
    }
}
