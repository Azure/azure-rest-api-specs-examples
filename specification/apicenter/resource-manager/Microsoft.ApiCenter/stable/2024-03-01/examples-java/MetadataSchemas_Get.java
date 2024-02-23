
/**
 * Samples for MetadataSchemas Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/MetadataSchemas_Get.json
     */
    /**
     * Sample code: MetadataSchemas_Get.
     * 
     * @param manager Entry point to ApiCenterManager.
     */
    public static void metadataSchemasGet(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        manager.metadataSchemas().getWithResponse("contoso-resources", "contoso", "lastName",
            com.azure.core.util.Context.NONE);
    }
}
