
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
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void replicationDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getReplications().delete("myResourceGroup", "myRegistry",
            "myReplication", com.azure.core.util.Context.NONE);
    }
}
