
/**
 * Samples for ServiceFabricSchedules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ServiceFabricSchedules_Get.json
     */
    /**
     * Sample code: ServiceFabricSchedules_Get.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceFabricSchedulesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.serviceFabricSchedules().getWithResponse("resourceGroupName", "{labName}", "@me", "{serviceFrabicName}",
            "{scheduleName}", null, com.azure.core.util.Context.NONE);
    }
}
