
/**
 * Samples for ContainerAppsRevisionReplicas ListReplicas.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/Replicas_List.json
     */
    /**
     * Sample code: List Container App's replicas.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listContainerAppSReplicas(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsRevisionReplicas().listReplicasWithResponse("workerapps-rg-xj", "myapp", "myapp--0wlqy09",
            com.azure.core.util.Context.NONE);
    }
}
