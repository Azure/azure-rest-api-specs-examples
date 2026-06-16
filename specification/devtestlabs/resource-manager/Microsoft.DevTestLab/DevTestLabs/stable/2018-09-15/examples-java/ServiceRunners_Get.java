
/**
 * Samples for ServiceRunners Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ServiceRunners_Get.json
     */
    /**
     * Sample code: ServiceRunners_Get.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void serviceRunnersGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.serviceRunners().getWithResponse("resourceGroupName", "{devtestlabName}", "{servicerunnerName}",
            com.azure.core.util.Context.NONE);
    }
}
