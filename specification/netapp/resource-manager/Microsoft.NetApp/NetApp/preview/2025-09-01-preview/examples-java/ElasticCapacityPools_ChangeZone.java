
import com.azure.resourcemanager.netapp.models.ChangeZoneRequest;

/**
 * Samples for ElasticCapacityPools ChangeZone.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticCapacityPools_ChangeZone.json
     */
    /**
     * Sample code: ElasticCapacityPools_ChangeZone.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticCapacityPoolsChangeZone(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticCapacityPools().changeZone("myRG", "account1", "pool1", new ChangeZoneRequest().withNewZone("3"),
            com.azure.core.util.Context.NONE);
    }
}
