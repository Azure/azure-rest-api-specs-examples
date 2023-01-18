/** Samples for Disks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Disks_Get.json
     */
    /**
     * Sample code: Disks_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void disksGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .disks()
            .getWithResponse(
                "resourceGroupName", "{labName}", "@me", "{diskName}", null, com.azure.core.util.Context.NONE);
    }
}
