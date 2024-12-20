
import com.azure.resourcemanager.storagepool.models.IscsiLun;
import com.azure.resourcemanager.storagepool.models.IscsiTargetAclMode;
import java.util.Arrays;

/**
 * Samples for IscsiTargets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Put.json
     */
    /**
     * Sample code: Create or Update iSCSI Target.
     * 
     * @param manager Entry point to StoragePoolManager.
     */
    public static void createOrUpdateISCSITarget(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.iscsiTargets().define("myIscsiTarget").withExistingDiskPool("myResourceGroup", "myDiskPool")
            .withAclMode(IscsiTargetAclMode.DYNAMIC).withTargetIqn("iqn.2005-03.org.iscsi:server1")
            .withLuns(Arrays.asList(new IscsiLun().withName("lun0").withManagedDiskAzureResourceId(
                "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1")))
            .create();
    }
}
