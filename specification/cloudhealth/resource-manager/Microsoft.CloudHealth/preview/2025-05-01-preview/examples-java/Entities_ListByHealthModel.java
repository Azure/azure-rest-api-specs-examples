
/**
 * Samples for Entities ListByHealthModel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Entities_ListByHealthModel.json
     */
    /**
     * Sample code: Entities_ListByHealthModel.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void entitiesListByHealthModel(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().listByHealthModel("rgopenapi", "gPWT6GP85xRV248L7LhNRTD--2Yc73wu-5Qk-0tS", null,
            com.azure.core.util.Context.NONE);
    }
}
