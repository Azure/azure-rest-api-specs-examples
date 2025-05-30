
/**
 * Samples for ContainerAppsRevisionReplicas GetReplica.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/Replicas_Get.json
     */
    /**
     * Sample code: Get Container App's revision replica.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getContainerAppSRevisionReplica(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsRevisionReplicas().getReplicaWithResponse("workerapps-rg-xj", "myapp", "myapp--0wlqy09",
            "myapp--0wlqy09-5d9774cff-5wnd8", com.azure.core.util.Context.NONE);
    }
}
