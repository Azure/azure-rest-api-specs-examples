
import com.azure.resourcemanager.oracledatabase.models.VirtualNetworkAddressProperties;

/**
 * Samples for VirtualNetworkAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/virtualNetworkAddresses_create.
     * json
     */
    /**
     * Sample code: Create Virtual Network Address.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        createVirtualNetworkAddress(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.virtualNetworkAddresses().define("hostname1").withExistingCloudVmCluster("rg000", "cluster1")
            .withProperties(
                new VirtualNetworkAddressProperties().withIpAddress("192.168.0.1").withVmOcid("ocid1..aaaa"))
            .create();
    }
}
