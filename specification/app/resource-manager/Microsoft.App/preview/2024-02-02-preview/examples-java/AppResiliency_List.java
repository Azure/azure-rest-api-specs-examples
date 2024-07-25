
/**
 * Samples for AppResiliency List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/AppResiliency_List.json
     */
    /**
     * Sample code: List App Resiliency.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listAppResiliency(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.appResiliencies().list("rg", "testcontainerApp0", com.azure.core.util.Context.NONE);
    }
}
