
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Usages_List.json
     */
    /**
     * Sample code: List usages.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listUsages(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.usages().list("westus", com.azure.core.util.Context.NONE);
    }
}
