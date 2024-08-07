
import com.azure.resourcemanager.oracledatabase.models.AddRemoveDbNode;
import java.util.Arrays;

/**
 * Samples for CloudVmClusters AddVms.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/vmClusters_addVms.json
     */
    /**
     * Sample code: Add VMs to VM Cluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void addVMsToVMCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().addVms("rg000", "cluster1",
            new AddRemoveDbNode().withDbServers(Arrays.asList("ocid1..aaaa", "ocid1..aaaaaa")),
            com.azure.core.util.Context.NONE);
    }
}
