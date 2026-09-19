
/**
 * Samples for ManagedEnvironmentsStorages Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironmentsStorages_Delete.json
     */
    /**
     * Sample code: List environments storages by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listEnvironmentsStoragesBySubscription(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentsStorages().deleteWithResponse("examplerg", "managedEnv", "jlaw-demo1",
            com.azure.core.util.Context.NONE);
    }
}
