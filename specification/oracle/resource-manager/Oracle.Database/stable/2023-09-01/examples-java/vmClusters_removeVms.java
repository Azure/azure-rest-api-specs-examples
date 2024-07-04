
import com.azure.resourcemanager.oracledatabase.models.AddRemoveDbNode;
import java.util.Arrays;

/**
 * Samples for CloudVmClusters RemoveVms.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/vmClusters_removeVms.json
     */
    /**
     * Sample code: Remove VMs from VM Cluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void removeVMsFromVMCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().removeVms("rg000", "cluster1",
            new AddRemoveDbNode().withDbServers(Arrays.asList("ocid1..aaaa")), com.azure.core.util.Context.NONE);
    }
}
