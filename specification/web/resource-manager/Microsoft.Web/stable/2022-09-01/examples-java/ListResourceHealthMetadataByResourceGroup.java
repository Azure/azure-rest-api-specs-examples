/** Samples for ResourceHealthMetadata ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/ListResourceHealthMetadataByResourceGroup.json
     */
    /**
     * Sample code: List ResourceHealthMetadata for a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listResourceHealthMetadataForAResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getResourceHealthMetadatas()
            .listByResourceGroup("Default-Web-NorthCentralUS", com.azure.core.util.Context.NONE);
    }
}
