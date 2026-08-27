
/**
 * Samples for HorizonDbAdministrators Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Administrators_Delete.json
     */
    /**
     * Sample code: Delete a HorizonDB administrator.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAHorizonDBAdministrator(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbAdministrators().delete("exampleresourcegroup", "examplecluster",
            "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", com.azure.core.util.Context.NONE);
    }
}
