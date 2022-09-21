import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.CapacityPool;

/** Samples for Pools Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-05-01/examples/Pools_Update.json
     */
    /**
     * Sample code: Pools_Update.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void poolsUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        CapacityPool resource = manager.pools().getWithResponse("myRG", "account1", "pool1", Context.NONE).getValue();
        resource.update().apply();
    }
}
