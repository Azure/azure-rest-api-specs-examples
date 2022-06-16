import com.azure.core.util.Context;

/** Samples for Replications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-12-01-preview/examples/ReplicationList.json
     */
    /**
     * Sample code: ReplicationList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void replicationList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .containerRegistries()
            .manager()
            .serviceClient()
            .getReplications()
            .list("myResourceGroup", "myRegistry", Context.NONE);
    }
}
