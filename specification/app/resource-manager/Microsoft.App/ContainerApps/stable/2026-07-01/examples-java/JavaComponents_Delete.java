
/**
 * Samples for JavaComponents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/JavaComponents_Delete.json
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
