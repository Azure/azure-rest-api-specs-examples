/** Samples for Replications Delete. */
public final class Main {
    /*
     * x-ms-original-file: mgmt_containerregistry_add_readonly/specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/stable/2023-07-01/examples/ReplicationDelete.json
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
            .delete("myResourceGroup", "myRegistry", "myReplication", com.azure.core.util.Context.NONE);
    }
}
