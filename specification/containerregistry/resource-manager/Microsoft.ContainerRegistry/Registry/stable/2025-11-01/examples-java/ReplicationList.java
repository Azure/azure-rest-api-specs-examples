
/**
 * Samples for Replications List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ReplicationList.json
     */
    /**
     * Sample code: ReplicationList.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void replicationList(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getReplications().list("myResourceGroup", "myRegistry",
            com.azure.core.util.Context.NONE);
    }
}
