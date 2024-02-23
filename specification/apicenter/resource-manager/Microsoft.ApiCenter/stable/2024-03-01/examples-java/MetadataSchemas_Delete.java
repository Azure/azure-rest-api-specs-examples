
/**
 * Samples for MetadataSchemas Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/MetadataSchemas_Delete.
     * json
     */
    /**
     * Sample code: MetadataSchemas_Delete.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void metadataSchemasDelete(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.metadataSchemas().deleteWithResponse("contoso-resources", "contoso", "author",
            com.azure.core.util.Context.NONE);
    }
}
