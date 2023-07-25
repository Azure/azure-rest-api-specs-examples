/** Samples for Pools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-11-01/examples/Pools_Get.json
     */
    /**
     * Sample code: Pools_Get.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void poolsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.pools().getWithResponse("myRG", "account1", "pool1", com.azure.core.util.Context.NONE);
    }
}
