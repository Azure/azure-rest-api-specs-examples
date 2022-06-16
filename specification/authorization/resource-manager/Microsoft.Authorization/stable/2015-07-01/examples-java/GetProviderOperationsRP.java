import com.azure.core.util.Context;

/** Samples for ProviderOperationsMetadata Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2015-07-01/examples/GetProviderOperationsRP.json
     */
    /**
     * Sample code: Get provider operations metadata for resource provider.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getProviderOperationsMetadataForResourceProvider(
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
