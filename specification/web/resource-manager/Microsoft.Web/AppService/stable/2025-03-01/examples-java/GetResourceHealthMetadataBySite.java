
/**
 * Samples for ResourceHealthMetadata GetBySite.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * GetResourceHealthMetadataBySite.json
     */
    /**
     * Sample code: Get ResourceHealthMetadata.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getResourceHealthMetadata(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getResourceHealthMetadatas().getBySiteWithResponse(
            "Default-Web-NorthCentralUS", "newsiteinnewASE-NCUS", com.azure.core.util.Context.NONE);
    }
}
