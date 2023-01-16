/** Samples for ArtifactSources Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/artifactsource_delete.json
     */
    /**
     * Sample code: Delete artifact source.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void deleteArtifactSource(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .artifactSources()
            .deleteByResourceGroupWithResponse("myResourceGroup", "myArtifactSource", com.azure.core.util.Context.NONE);
    }
}
