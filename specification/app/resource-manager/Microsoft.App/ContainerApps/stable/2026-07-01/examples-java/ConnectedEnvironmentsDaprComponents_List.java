
/**
 * Samples for ConnectedEnvironmentsDaprComponents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ConnectedEnvironmentsDaprComponents_List.json
     */
    /**
     * Sample code: List Dapr Components.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listDaprComponents(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironmentsDaprComponents().list("examplerg", "myenvironment",
            com.azure.core.util.Context.NONE);
    }
}
