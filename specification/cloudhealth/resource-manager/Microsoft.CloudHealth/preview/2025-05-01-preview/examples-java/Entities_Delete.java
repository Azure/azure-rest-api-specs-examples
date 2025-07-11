
/**
 * Samples for Entities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/Entities_Delete.json
     */
    /**
     * Sample code: Entities_Delete.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void entitiesDelete(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.entities().deleteWithResponse("rgopenapi", "model1", "U4VTRFlUkm9kR6H23-c-6U-XHq7n",
            com.azure.core.util.Context.NONE);
    }
}
