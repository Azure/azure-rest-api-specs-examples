
/**
 * Samples for Builders List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/Builders_ListBySubscription.
     * json
     */
    /**
     * Sample code: Builders_ListBySubscription_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        buildersListBySubscription0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.builders().list(com.azure.core.util.Context.NONE);
    }
}
