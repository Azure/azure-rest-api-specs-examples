
/**
 * Samples for ArtifactSources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ArtifactSources_List.json
     */
    /**
     * Sample code: ArtifactSources_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void artifactSourcesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.artifactSources().list("resourceGroupName", "{labName}", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
