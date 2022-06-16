import com.azure.core.util.Context;

/** Samples for ProviderOperationsMetadata List. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-01-01-preview/examples/GetAllProviderOperations.json
     */
    /**
     * Sample code: List provider operations metadata for all resource providers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listProviderOperationsMetadataForAllResourceProviders(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getProviderOperationsMetadatas()
            .list(null, Context.NONE);
    }
}
