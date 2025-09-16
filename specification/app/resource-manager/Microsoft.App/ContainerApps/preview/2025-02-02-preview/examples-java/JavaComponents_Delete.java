
/**
 * Samples for JavaComponents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * JavaComponents_Delete.json
     */
    /**
     * Sample code: Delete Java Component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteJavaComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.javaComponents().delete("examplerg", "myenvironment", "myjavacomponent",
            com.azure.core.util.Context.NONE);
    }
}
