
/**
 * Samples for DaprSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/DaprSubscriptions_List.json
     */
    /**
     * Sample code: List Dapr subscriptions.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listDaprSubscriptions(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprSubscriptions().list("examplerg", "myenvironment", com.azure.core.util.Context.NONE);
    }
}
