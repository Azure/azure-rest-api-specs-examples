
/**
 * Samples for Relationships Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/Relationships_Delete.json
     */
    /**
     * Sample code: Relationships_Delete.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void relationshipsDelete(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.relationships().delete("rgopenapi", "model1", "rel1", com.azure.core.util.Context.NONE);
    }
}
