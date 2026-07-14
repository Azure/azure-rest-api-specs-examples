
/**
 * Samples for Entities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Entities_Delete.json
     */
    /**
     * Sample code: Entities_Delete.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void entitiesDelete(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().delete("online-store-rg", "online-store", "catalog-storage",
            com.azure.core.util.Context.NONE);
    }
}
