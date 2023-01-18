/** Samples for Environments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Environments_List.json
     */
    /**
     * Sample code: Environments_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void environmentsList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .environments()
            .list("resourceGroupName", "{labName}", "@me", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
