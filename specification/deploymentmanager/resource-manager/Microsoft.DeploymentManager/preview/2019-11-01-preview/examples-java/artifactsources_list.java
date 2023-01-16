/** Samples for ArtifactSources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/artifactsources_list.json
     */
    /**
     * Sample code: List steps.
     *
     * @param manager Entry point to DeploymentManager.
     */
    public static void listSteps(com.azure.resourcemanager.deploymentmanager.DeploymentManager manager) {
        manager.artifactSources().listWithResponse("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
