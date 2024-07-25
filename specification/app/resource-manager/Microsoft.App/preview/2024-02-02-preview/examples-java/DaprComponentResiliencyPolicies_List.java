
/**
 * Samples for DaprComponentResiliencyPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/
     * DaprComponentResiliencyPolicies_List.json
     */
    /**
     * Sample code: List Dapr component resiliency policies.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listDaprComponentResiliencyPolicies(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.daprComponentResiliencyPolicies().list("examplerg", "myenvironment", "mydaprcomponent",
            com.azure.core.util.Context.NONE);
    }
}
