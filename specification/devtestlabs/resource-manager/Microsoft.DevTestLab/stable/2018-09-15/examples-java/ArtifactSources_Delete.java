/** Samples for ArtifactSources Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_Delete.json
     */
    /**
     * Sample code: ArtifactSources_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void artifactSourcesDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .artifactSources()
            .deleteWithResponse(
                "resourceGroupName", "{labName}", "{artifactSourceName}", com.azure.core.util.Context.NONE);
    }
}
