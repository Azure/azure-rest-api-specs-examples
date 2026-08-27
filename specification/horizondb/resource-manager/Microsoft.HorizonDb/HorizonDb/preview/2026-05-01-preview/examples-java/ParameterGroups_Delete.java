
/**
 * Samples for HorizonDbParameterGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ParameterGroups_Delete.json
     */
    /**
     * Sample code: Delete a HorizonDB parameter group.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAHorizonDBParameterGroup(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbParameterGroups().delete("exampleresourcegroup", "exampleparametergroup",
            com.azure.core.util.Context.NONE);
    }
}
