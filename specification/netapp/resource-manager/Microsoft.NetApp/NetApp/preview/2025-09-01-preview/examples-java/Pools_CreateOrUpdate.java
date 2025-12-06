
import com.azure.resourcemanager.netapp.models.QosType;
import com.azure.resourcemanager.netapp.models.ServiceLevel;

/**
 * Samples for Pools CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Pools_CreateOrUpdate.json
     */
    /**
     * Sample code: Pools_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void poolsCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.pools().define("pool1").withRegion("eastus").withExistingNetAppAccount("myRG", "account1")
            .withSize(4398046511104L).withServiceLevel(ServiceLevel.PREMIUM).withQosType(QosType.AUTO).create();
    }
}
