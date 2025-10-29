
/**
 * Samples for JavaComponents List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/ContainerApps/stable/2025-07-01/examples/
     * JavaComponents_List_ServiceBind.json
     */
    /**
     * Sample code: List Java Components with ServiceBinds.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listJavaComponentsWithServiceBinds(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.javaComponents().list("examplerg", "myenvironment", com.azure.core.util.Context.NONE);
    }
}
