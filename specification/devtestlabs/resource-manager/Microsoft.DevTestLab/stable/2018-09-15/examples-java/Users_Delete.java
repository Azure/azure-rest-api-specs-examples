/** Samples for Users Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Users_Delete.json
     */
    /**
     * Sample code: Users_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void usersDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.users().delete("resourceGroupName", "{devtestlabName}", "{userName}", com.azure.core.util.Context.NONE);
    }
}
