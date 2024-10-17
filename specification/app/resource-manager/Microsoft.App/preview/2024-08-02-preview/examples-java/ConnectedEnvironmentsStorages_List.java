
/**
 * Samples for ConnectedEnvironmentsStorages List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/
     * ConnectedEnvironmentsStorages_List.json
     */
    /**
     * Sample code: List environments storages by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listEnvironmentsStoragesBySubscription(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironmentsStorages().listWithResponse("examplerg", "managedEnv",
            com.azure.core.util.Context.NONE);
    }
}
