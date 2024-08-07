
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Usages_List.json
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
