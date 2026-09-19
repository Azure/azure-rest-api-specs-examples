
/**
 * Samples for ManagedEnvironmentsStorages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedEnvironmentsStorages_List.json
     */
    /**
     * Sample code: List environments storages by subscription.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listEnvironmentsStoragesBySubscription(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentsStorages().listWithResponse("examplerg", "managedEnv",
            com.azure.core.util.Context.NONE);
    }
}
