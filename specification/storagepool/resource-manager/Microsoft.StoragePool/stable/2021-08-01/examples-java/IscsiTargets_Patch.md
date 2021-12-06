Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-storagepool_1.0.0-beta.1/sdk/storagepool/azure-resourcemanager-storagepool/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storagepool.models.Acl;
import com.azure.resourcemanager.storagepool.models.IscsiLun;
import com.azure.resourcemanager.storagepool.models.IscsiTarget;
import java.util.Arrays;

/** Samples for IscsiTargets Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/IscsiTargets_Patch.json
     */
    /**
     * Sample code: Update iSCSI Target.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void updateISCSITarget(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        IscsiTarget resource =
            manager
                .iscsiTargets()
                .getWithResponse("myResourceGroup", "myDiskPool", "myIscsiTarget", Context.NONE)
                .getValue();
        resource
            .update()
            .withStaticAcls(
                Arrays
                    .asList(
                        new Acl()
                            .withInitiatorIqn("iqn.2005-03.org.iscsi:client")
                            .withMappedLuns(Arrays.asList("lun0"))))
            .withLuns(
                Arrays
                    .asList(
                        new IscsiLun()
                            .withName("lun0")
                            .withManagedDiskAzureResourceId(
                                "/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1")))
            .apply();
    }
}
```
