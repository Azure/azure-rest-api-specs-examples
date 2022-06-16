import com.azure.core.util.Context;

/** Samples for ManagedEnvironments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/ManagedEnvironments_ListBySubscription.json
     */
    /**
     * Sample code: List environments by subscription.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listEnvironmentsBySubscription(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedEnvironments().list(Context.NONE);
    }
}
