
/**
 * Samples for HorizonDbParameterGroups ListConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/ParameterGroups_ListConnections.json
     */
    /**
     * Sample code: List connections for a HorizonDB parameter group.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listConnectionsForAHorizonDBParameterGroup(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbParameterGroups().listConnections("exampleresourcegroup", "exampleparametergroup",
            com.azure.core.util.Context.NONE);
    }
}
