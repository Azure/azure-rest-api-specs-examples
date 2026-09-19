
/**
 * Samples for JavaComponents List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/JavaComponents_List.json
     */
    /**
     * Sample code: List Java Components.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listJavaComponents(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.javaComponents().list("examplerg", "myenvironment", com.azure.core.util.Context.NONE);
    }
}
