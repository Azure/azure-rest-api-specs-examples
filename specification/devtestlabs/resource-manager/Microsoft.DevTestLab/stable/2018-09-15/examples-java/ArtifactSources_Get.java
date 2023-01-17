/** Samples for ArtifactSources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_Get.json
     */
    /**
     * Sample code: ArtifactSources_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void artifactSourcesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .artifactSources()
            .getWithResponse(
                "resourceGroupName", "{labName}", "{artifactSourceName}", null, com.azure.core.util.Context.NONE);
    }
}
