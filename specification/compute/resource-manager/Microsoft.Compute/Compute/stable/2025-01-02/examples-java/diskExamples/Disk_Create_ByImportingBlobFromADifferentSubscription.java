
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_Create_ByImportingBlobFromADifferentSubscription.json
     */
    /**
     * Sample code: create a managed disk by importing an unmanaged blob from a different subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createAManagedDiskByImportingAnUnmanagedBlobFromADifferentSubscription(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk", new DiskInner()
            .withLocation("West US")
            .withCreationData(new CreationData().withCreateOption(DiskCreateOption.IMPORT).withStorageAccountId(
                "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount")
                .withSourceUri("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd")),
            com.azure.core.util.Context.NONE);
    }
}
