/** Samples for Metadata GetEntity. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Metadata_GetEntity.json
     */
    /**
     * Sample code: GetMetadata.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void getMetadata(com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.metadatas().getEntityWithResponse("status", com.azure.core.util.Context.NONE);
    }
}
