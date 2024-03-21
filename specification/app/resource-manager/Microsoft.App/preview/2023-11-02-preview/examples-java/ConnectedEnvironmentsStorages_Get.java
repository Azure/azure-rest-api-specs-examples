
/**
 * Samples for ConnectedEnvironmentsStorages Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/
     * ConnectedEnvironmentsStorages_Get.json
     */
    /**
     * Sample code: get a environments storage properties by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getAEnvironmentsStoragePropertiesBySubscription(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironmentsStorages().getWithResponse("examplerg", "env", "jlaw-demo1",
            com.azure.core.util.Context.NONE);
    }
}
