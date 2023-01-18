/** Samples for ServiceFabricSchedules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ServiceFabricSchedules_Get.json
     */
    /**
     * Sample code: ServiceFabricSchedules_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricSchedulesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .serviceFabricSchedules()
            .getWithResponse(
                "resourceGroupName",
                "{labName}",
                "@me",
                "{serviceFrabicName}",
                "{scheduleName}",
                null,
                com.azure.core.util.Context.NONE);
    }
}
