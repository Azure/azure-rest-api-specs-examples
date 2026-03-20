
/**
 * Samples for Replications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ReplicationDelete.json
     */
    /**
     * Sample code: ReplicationDelete.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void replicationDelete(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getReplications().delete("myResourceGroup", "myRegistry", "myReplication",
            com.azure.core.util.Context.NONE);
    }
}
