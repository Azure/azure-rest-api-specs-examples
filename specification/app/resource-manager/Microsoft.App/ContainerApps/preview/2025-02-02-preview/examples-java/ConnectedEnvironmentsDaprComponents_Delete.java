
/**
 * Samples for ConnectedEnvironmentsDaprComponents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * ConnectedEnvironmentsDaprComponents_Delete.json
     */
    /**
     * Sample code: Delete dapr component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteDaprComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironmentsDaprComponents().delete("examplerg", "myenvironment", "reddog",
            com.azure.core.util.Context.NONE);
    }
}
