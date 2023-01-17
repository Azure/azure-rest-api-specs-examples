/** Samples for Artifacts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Artifacts_Get.json
     */
    /**
     * Sample code: Artifacts_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void artifactsGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .artifacts()
            .getWithResponse(
                "resourceGroupName",
                "{labName}",
                "{artifactSourceName}",
                "{artifactName}",
                null,
                com.azure.core.util.Context.NONE);
    }
}
