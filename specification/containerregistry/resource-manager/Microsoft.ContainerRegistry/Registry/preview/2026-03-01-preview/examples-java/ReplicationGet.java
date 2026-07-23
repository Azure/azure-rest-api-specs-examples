
/**
 * Samples for Replications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ReplicationGet.json
     */
    /**
     * Sample code: ReplicationGet.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void replicationGet(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getReplications().getWithResponse("myResourceGroup", "myRegistry", "myReplication",
            com.azure.core.util.Context.NONE);
    }
}
