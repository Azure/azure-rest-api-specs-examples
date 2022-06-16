import com.azure.core.util.Context;

/** Samples for ProviderOperationsMetadata Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/preview/2018-01-01-preview/examples/GetProviderOperationsRP.json
     */
    /**
     * Sample code: List provider operations metadata for resource provider.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listProviderOperationsMetadataForResourceProvider(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getProviderOperationsMetadatas()
            .getWithResponse("resourceProviderNamespace", null, Context.NONE);
    }
}
