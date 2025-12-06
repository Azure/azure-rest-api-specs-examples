
import com.azure.resourcemanager.netapp.models.QosType;
import com.azure.resourcemanager.netapp.models.ServiceLevel;

/**
 * Samples for Pools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Pools_CreateOrUpdate_CustomThroughput.json
     */
    /**
     * Sample code: Pools_CreateOrUpdate_CustomThroughput.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        poolsCreateOrUpdateCustomThroughput(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.pools().define("customPool1").withRegion("eastus").withExistingNetAppAccount("myRG", "account1")
            .withSize(4398046511104L).withServiceLevel(ServiceLevel.FLEXIBLE).withCustomThroughputMibps(128)
            .withQosType(QosType.MANUAL).create();
    }
}
