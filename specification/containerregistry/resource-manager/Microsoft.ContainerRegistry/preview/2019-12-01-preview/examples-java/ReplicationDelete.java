import com.azure.core.util.Context;

/** Samples for Replications Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/ReplicationDelete.json
     */
    /**
     * Sample code: ReplicationDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void replicationDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getReplications()
            .delete("myResourceGroup", "myRegistry", "myReplication", Context.NONE);
    }
}
