
/**
 * Samples for AppResiliency Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/AppResiliency_Delete.json
     */
    /**
     * Sample code: Delete App Resiliency.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteAppResiliency(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.appResiliencies().deleteWithResponse("rg", "testcontainerApp0", "resiliency-policy-1",
            com.azure.core.util.Context.NONE);
    }
}
