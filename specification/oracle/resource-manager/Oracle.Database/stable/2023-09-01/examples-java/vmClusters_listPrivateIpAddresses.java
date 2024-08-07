
import com.azure.resourcemanager.oracledatabase.models.PrivateIpAddressesFilter;

/**
 * Samples for CloudVmClusters ListPrivateIpAddresses.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * vmClusters_listPrivateIpAddresses.json
     */
    /**
     * Sample code: List Private IP Addresses for VM Cluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listPrivateIPAddressesForVMCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().listPrivateIpAddressesWithResponse("rg000", "cluster1",
            new PrivateIpAddressesFilter().withSubnetId("ocid1..aaaaaa").withVnicId("ocid1..aaaaa"),
            com.azure.core.util.Context.NONE);
    }
}
