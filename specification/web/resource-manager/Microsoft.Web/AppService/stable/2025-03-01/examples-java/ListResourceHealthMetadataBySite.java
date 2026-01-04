
/**
 * Samples for ResourceHealthMetadata ListBySiteSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * ListResourceHealthMetadataBySite.json
     */
    /**
     * Sample code: List ResourceHealthMetadata for a site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listResourceHealthMetadataForASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getResourceHealthMetadatas().listBySiteSlot(
            "Default-Web-NorthCentralUS", "newsiteinnewASE-NCUS", "Production", com.azure.core.util.Context.NONE);
    }
}
