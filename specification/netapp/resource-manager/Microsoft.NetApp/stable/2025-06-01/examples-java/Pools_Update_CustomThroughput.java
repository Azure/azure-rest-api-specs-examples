
import com.azure.resourcemanager.netapp.models.CapacityPool;

/**
 * Samples for Pools Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2025-06-01/examples/Pools_Update_CustomThroughput.
     * json
     */
    /**
     * Sample code: Pools_Update_CustomThroughput.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void poolsUpdateCustomThroughput(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        CapacityPool resource = manager.pools()
            .getWithResponse("myRG", "account1", "customPool1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
