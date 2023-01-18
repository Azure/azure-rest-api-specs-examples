/** Samples for ServiceFabrics Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabrics_Get.json
     */
    /**
     * Sample code: ServiceFabrics_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricsGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .serviceFabrics()
            .getWithResponse(
                "resourceGroupName",
                "{labName}",
                "{userName}",
                "{serviceFabricName}",
                null,
                com.azure.core.util.Context.NONE);
    }
}
