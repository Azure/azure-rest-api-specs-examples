
/**
 * Samples for Artifacts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Artifacts_List.json
     */
    /**
     * Sample code: Artifacts_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void artifactsList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.artifacts().list("resourceGroupName", "{labName}", "{artifactSourceName}", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
