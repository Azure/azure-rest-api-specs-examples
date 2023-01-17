/** Samples for Labs Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Labs_Delete.json
     */
    /**
     * Sample code: Labs_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void labsDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.labs().delete("resourceGroupName", "{labName}", com.azure.core.util.Context.NONE);
    }
}
