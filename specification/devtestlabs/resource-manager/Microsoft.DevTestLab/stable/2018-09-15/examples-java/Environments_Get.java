/** Samples for Environments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Environments_Get.json
     */
    /**
     * Sample code: Environments_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void environmentsGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .environments()
            .getWithResponse(
                "resourceGroupName", "{labName}", "@me", "{environmentName}", null, com.azure.core.util.Context.NONE);
    }
}
