
import com.azure.resourcemanager.oracledatabase.models.CloudVmCluster;

/**
 * Samples for CloudVmClusters Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/vmClusters_patch.json
     */
    /**
     * Sample code: Patch VM Cluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void patchVMCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        CloudVmCluster resource = manager.cloudVmClusters()
            .getByResourceGroupWithResponse("rg000", "cluster1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
