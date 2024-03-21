
/**
 * Samples for JavaComponents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/JavaComponents_Delete.json
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
