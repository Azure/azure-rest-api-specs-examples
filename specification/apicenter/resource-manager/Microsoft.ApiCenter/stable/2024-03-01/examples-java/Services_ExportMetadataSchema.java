
/**
 * Samples for Services ExportMetadataSchema.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/
     * Services_ExportMetadataSchema.json
     */
    /**
     * Sample code: Services_ExportMetadataSchema.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void servicesExportMetadataSchema(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.services().exportMetadataSchema("contoso-resources", "contoso", null, com.azure.core.util.Context.NONE);
    }
}
