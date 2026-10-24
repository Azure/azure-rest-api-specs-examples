
/**
 * Samples for ParameterGroups ListConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ParameterGroups_ListConnections.json
     */
    /**
     * Sample code: List HorizonDB pools connected to a parameter group.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listHorizonDBPoolsConnectedToAParameterGroup(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.parameterGroups().listConnections("example-resource-group", "production-parameters",
            com.azure.core.util.Context.NONE);
    }
}
