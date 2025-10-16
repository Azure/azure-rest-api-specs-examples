
import com.azure.resourcemanager.oracledatabase.models.AddRemoveDbNode;
import java.util.Arrays;

/**
 * Samples for CloudVmClusters AddVms.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/vmClusters_addVms.json
     */
    /**
     * Sample code: CloudVmClusters_addVms.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void cloudVmClustersAddVms(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().addVms("rg000", "cluster1",
            new AddRemoveDbNode().withDbServers(Arrays.asList("ocid1..aaaa", "ocid1..aaaaaa")),
            com.azure.core.util.Context.NONE);
    }
}
