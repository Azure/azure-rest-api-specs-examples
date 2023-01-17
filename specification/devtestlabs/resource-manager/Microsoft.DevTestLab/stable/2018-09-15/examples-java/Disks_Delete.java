/** Samples for Disks Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Disks_Delete.json
     */
    /**
     * Sample code: Disks_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void disksDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .disks()
            .delete("resourceGroupName", "{labName}", "{userId}", "{diskName}", com.azure.core.util.Context.NONE);
    }
}
