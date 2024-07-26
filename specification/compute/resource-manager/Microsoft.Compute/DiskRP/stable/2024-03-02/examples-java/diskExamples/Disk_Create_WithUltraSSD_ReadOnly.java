
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.DiskSku;
import com.azure.resourcemanager.compute.models.DiskStorageAccountTypes;
import com.azure.resourcemanager.compute.models.Encryption;
import com.azure.resourcemanager.compute.models.EncryptionType;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/
     * Disk_Create_WithUltraSSD_ReadOnly.json
     */
    /**
     * Sample code: Create a managed disk with ultra account type with readOnly property set.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAManagedDiskWithUltraAccountTypeWithReadOnlyPropertySet(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().createOrUpdate("myResourceGroup",
            "myUltraReadOnlyDisk",
            new DiskInner().withLocation("West US")
                .withSku(new DiskSku().withName(DiskStorageAccountTypes.ULTRA_SSD_LRS))
                .withCreationData(
                    new CreationData().withCreateOption(DiskCreateOption.EMPTY).withLogicalSectorSize(4096))
                .withDiskSizeGB(200).withDiskIopsReadWrite(125L).withDiskMBpsReadWrite(3000L)
                .withEncryption(new Encryption().withType(EncryptionType.ENCRYPTION_AT_REST_WITH_PLATFORM_KEY)),
            com.azure.core.util.Context.NONE);
    }
}
