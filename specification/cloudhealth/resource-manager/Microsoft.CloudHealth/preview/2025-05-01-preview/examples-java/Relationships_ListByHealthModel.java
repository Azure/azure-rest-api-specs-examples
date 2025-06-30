
/**
 * Samples for Relationships ListByHealthModel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Relationships_ListByHealthModel.json
     */
    /**
     * Sample code: Relationships_ListByHealthModel.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        relationshipsListByHealthModel(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.relationships().listByHealthModel("rgopenapi", "model1", null, com.azure.core.util.Context.NONE);
    }
}
