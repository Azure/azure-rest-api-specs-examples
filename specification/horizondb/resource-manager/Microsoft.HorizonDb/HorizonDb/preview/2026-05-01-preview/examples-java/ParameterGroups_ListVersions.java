
/**
 * Samples for HorizonDbParameterGroups ListVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ParameterGroups_ListVersions.json
     */
    /**
     * Sample code: List parameter groups filtered by version.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listParameterGroupsFilteredByVersion(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbParameterGroups().listVersions("exampleresourcegroup", "exampleparametergroup", 22,
            com.azure.core.util.Context.NONE);
    }
}
