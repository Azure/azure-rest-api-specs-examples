import com.azure.core.util.Context;

/** Samples for ManagedEnvironmentsStorages Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ManagedEnvironmentsStorages_Delete.json
     */
    /**
     * Sample code: List environments storages by subscription.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listEnvironmentsStoragesBySubscription(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironmentsStorages().deleteWithResponse("examplerg", "managedEnv", "jlaw-demo1", Context.NONE);
    }
}
