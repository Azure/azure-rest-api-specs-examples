
/**
 * Samples for DotNetComponents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/DotNetComponents_Delete.json
     */
    /**
     * Sample code: Delete .NET Component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteNETComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.dotNetComponents().delete("examplerg", "myenvironment", "mydotnetcomponent",
            com.azure.core.util.Context.NONE);
    }
}
