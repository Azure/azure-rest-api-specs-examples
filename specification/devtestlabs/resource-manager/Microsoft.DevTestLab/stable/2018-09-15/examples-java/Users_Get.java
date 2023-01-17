/** Samples for Users Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Users_Get.json
     */
    /**
     * Sample code: Users_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void usersGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .users()
            .getWithResponse(
                "resourceGroupName", "{devtestlabName}", "{userName}", null, com.azure.core.util.Context.NONE);
    }
}
