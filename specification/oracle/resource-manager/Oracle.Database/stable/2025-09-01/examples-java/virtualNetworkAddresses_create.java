
import com.azure.resourcemanager.oracledatabase.models.VirtualNetworkAddressProperties;

/**
 * Samples for VirtualNetworkAddresses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/virtualNetworkAddresses_create.json
     */
    /**
     * Sample code: VirtualNetworkAddresses_createOrUpdate.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        virtualNetworkAddressesCreateOrUpdate(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.virtualNetworkAddresses().define("hostname1").withExistingCloudVmCluster("rg000", "cluster1")
            .withProperties(
                new VirtualNetworkAddressProperties().withIpAddress("192.168.0.1").withVmOcid("ocid1..aaaa"))
            .create();
    }
}
