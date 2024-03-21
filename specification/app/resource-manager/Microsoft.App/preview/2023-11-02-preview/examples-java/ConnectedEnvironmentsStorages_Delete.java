
/**
 * Samples for ConnectedEnvironmentsStorages Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/
     * ConnectedEnvironmentsStorages_Delete.json
     */
    /**
     * Sample code: List environments storages by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listEnvironmentsStoragesBySubscription(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironmentsStorages().deleteWithResponse("examplerg", "env", "jlaw-demo1",
            com.azure.core.util.Context.NONE);
    }
}
