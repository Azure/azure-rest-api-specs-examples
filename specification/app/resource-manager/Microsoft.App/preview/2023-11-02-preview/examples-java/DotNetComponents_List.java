
/**
 * Samples for DotNetComponents List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/DotNetComponents_List.json
     */
    /**
     * Sample code: List .NET Components.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listNETComponents(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.dotNetComponents().list("examplerg", "myenvironment", com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/
     * DotNetComponents_List_ServiceBind.json
     */
    /**
     * Sample code: List .NET Components with ServiceBinds.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listNETComponentsWithServiceBinds(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.dotNetComponents().list("examplerg", "myenvironment", com.azure.core.util.Context.NONE);
    }
}
