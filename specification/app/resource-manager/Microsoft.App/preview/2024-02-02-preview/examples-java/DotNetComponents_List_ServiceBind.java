
/**
 * Samples for DotNetComponents List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
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
