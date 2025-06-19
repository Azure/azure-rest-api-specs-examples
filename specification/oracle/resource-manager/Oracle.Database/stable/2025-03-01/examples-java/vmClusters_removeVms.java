
import com.azure.resourcemanager.oracledatabase.models.AddRemoveDbNode;
import java.util.Arrays;

/**
 * Samples for CloudVmClusters RemoveVms.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/vmClusters_removeVms.json
     */
    /**
     * Sample code: CloudVmClusters_removeVms.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        cloudVmClustersRemoveVms(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().removeVms("rg000", "cluster1",
            new AddRemoveDbNode().withDbServers(Arrays.asList("ocid1..aaaa")), com.azure.core.util.Context.NONE);
    }
}
