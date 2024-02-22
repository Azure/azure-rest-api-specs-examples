
/**
 * Samples for MetadataSchemas Head.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/MetadataSchemas_Head.json
     */
    /**
     * Sample code: MetadataSchemas_Head.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void metadataSchemasHead(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.metadataSchemas().headWithResponse("contoso-resources", "contoso", "author",
            com.azure.core.util.Context.NONE);
    }
}
