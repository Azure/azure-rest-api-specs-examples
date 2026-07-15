
/**
 * Samples for Relationships ListByHealthModel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Relationships_ListByHealthModel.json
     */
    /**
     * Sample code: Relationships_ListByHealthModel.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        relationshipsListByHealthModel(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.relationships().listByHealthModel("online-store-rg", "online-store", null,
            com.azure.core.util.Context.NONE);
    }
}
