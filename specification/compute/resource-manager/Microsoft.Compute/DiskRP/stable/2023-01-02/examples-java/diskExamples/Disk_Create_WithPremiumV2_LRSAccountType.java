import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.DiskSku;
import com.azure.resourcemanager.compute.models.DiskStorageAccountTypes;

/** Samples for Disks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-01-02/examples/diskExamples/Disk_Create_WithPremiumV2_LRSAccountType.json
     */
    /**
     * Sample code: Create a managed disk with premium v2 account type.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAManagedDiskWithPremiumV2AccountType(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDisks()
            .createOrUpdate(
                "myResourceGroup",
                "myPremiumV2Disk",
                new DiskInner()
                    .withLocation("West US")
                    .withSku(new DiskSku().withName(DiskStorageAccountTypes.PREMIUM_V2_LRS))
                    .withCreationData(new CreationData().withCreateOption(DiskCreateOption.EMPTY))
                    .withDiskSizeGB(200)
                    .withDiskIopsReadWrite(125L)
                    .withDiskMBpsReadWrite(3000L),
                com.azure.core.util.Context.NONE);
    }
}
