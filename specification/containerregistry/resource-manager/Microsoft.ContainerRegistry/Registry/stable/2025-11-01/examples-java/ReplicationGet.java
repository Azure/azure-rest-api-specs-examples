
/**
 * Samples for Replications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/ReplicationGet.json
     */
    /**
     * Sample code: ReplicationGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void replicationGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getReplications().getWithResponse("myResourceGroup",
            "myRegistry", "myReplication", com.azure.core.util.Context.NONE);
    }
}
