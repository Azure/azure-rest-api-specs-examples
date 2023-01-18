/** Samples for ArtifactSources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_List.json
     */
    /**
     * Sample code: ArtifactSources_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void artifactSourcesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .artifactSources()
            .list("resourceGroupName", "{labName}", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
