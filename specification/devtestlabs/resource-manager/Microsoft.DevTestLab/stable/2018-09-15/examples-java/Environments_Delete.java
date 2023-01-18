/** Samples for Environments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Environments_Delete.json
     */
    /**
     * Sample code: Environments_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void environmentsDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .environments()
            .delete("resourceGroupName", "{labName}", "@me", "{environmentName}", com.azure.core.util.Context.NONE);
    }
}
