
/**
 * Samples for Entities ListByHealthModel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/Entities_ListByHealthModel.json
     */
    /**
     * Sample code: Entities_ListByHealthModel.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void entitiesListByHealthModel(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().listByHealthModel("online-store-rg", "online-store", null, com.azure.core.util.Context.NONE);
    }
}
