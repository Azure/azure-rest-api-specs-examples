import com.azure.core.util.Context;

/** Samples for Pools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-08-01/examples/Pools_Delete.json
     */
    /**
     * Sample code: Pools_Delete.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void poolsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.pools().delete("myRG", "account1", "pool1", Context.NONE);
    }
}
