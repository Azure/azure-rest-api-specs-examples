/** Samples for ProviderOperationsMetadata Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/examples/GetProviderOperationsRP.json
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
            .getWithResponse("resourceProviderNamespace", null, com.azure.core.util.Context.NONE);
    }
}
