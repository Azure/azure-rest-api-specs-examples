
/**
 * Samples for DaprComponents List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * DaprComponents_List.json
     */
    /**
     * Sample code: List Dapr Components.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listDaprComponents(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponents().list("examplerg", "myenvironment", com.azure.core.util.Context.NONE);
    }
}
