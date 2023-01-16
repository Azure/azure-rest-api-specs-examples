/** Samples for ArtifactSources GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/artifactsource_get.json
     */
    /**
     * Sample code: Get artifact source.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void getArtifactSource(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager
            .artifactSources()
            .getByResourceGroupWithResponse("myResourceGroup", "myArtifactSource", com.azure.core.util.Context.NONE);
    }
}
