
/**
 * Samples for Replications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/ReplicationDelete.json
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
