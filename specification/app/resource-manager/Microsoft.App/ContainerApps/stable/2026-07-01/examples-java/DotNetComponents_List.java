
/**
 * Samples for DotNetComponents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/DotNetComponents_List.json
     */
    /**
     * Sample code: List .NET Components.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listNETComponents(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.dotNetComponents().list("examplerg", "myenvironment", com.azure.core.util.Context.NONE);
    }
}
