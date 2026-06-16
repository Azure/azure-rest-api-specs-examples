
/**
 * Samples for ArtifactSources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/ArtifactSources_Get.json
     */
    /**
     * Sample code: ArtifactSources_Get.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void artifactSourcesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.artifactSources().getWithResponse("resourceGroupName", "{labName}", "{artifactSourceName}", null,
            com.azure.core.util.Context.NONE);
    }
}
