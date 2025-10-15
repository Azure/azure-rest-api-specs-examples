
import com.azure.resourcemanager.oracledatabase.models.PrivateIpAddressesFilter;

/**
 * Samples for CloudVmClusters ListPrivateIpAddresses.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/vmClusters_listPrivateIpAddresses.json
     */
    /**
     * Sample code: CloudVmClusters_listPrivateIpAddresses.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        cloudVmClustersListPrivateIpAddresses(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudVmClusters().listPrivateIpAddressesWithResponse("rg000", "cluster1",
            new PrivateIpAddressesFilter().withSubnetId("ocid1..aaaaaa").withVnicId("ocid1..aaaaa"),
            com.azure.core.util.Context.NONE);
    }
}
