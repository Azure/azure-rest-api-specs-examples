
import com.azure.resourcemanager.avs.models.NetAppVolume;

/**
 * Samples for Datastores CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Datastores_CreateOrUpdate.json
     */
    /**
     * Sample code: Datastores_CreateOrUpdate.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void datastoresCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.datastores().define("datastore1").withExistingCluster("group1", "cloud1", "cluster1")
            .withNetAppVolume(new NetAppVolume().withId(
                "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/ResourceGroup1/providers/Microsoft.NetApp/netAppAccounts/NetAppAccount1/capacityPools/CapacityPool1/volumes/NFSVol1"))
            .create();
    }
}
