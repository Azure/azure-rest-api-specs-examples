
/**
 * Samples for MetadataSchemas List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/MetadataSchemas_List.json
     */
    /**
     * Sample code: MetadataSchemas_ListByService.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void metadataSchemasListByService(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.metadataSchemas().list("contoso-resources", "contoso", null, com.azure.core.util.Context.NONE);
    }
}
