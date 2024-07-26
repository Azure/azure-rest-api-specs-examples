
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.DiskSecurityProfile;
import com.azure.resourcemanager.compute.models.DiskSecurityTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/
     * Disk_Create_FromImportSecure.json
     */
    /**
     * Sample code: Create a managed disk from ImportSecure create option.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createAManagedDiskFromImportSecureCreateOption(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("West US").withOsType(OperatingSystemTypes.WINDOWS).withCreationData(
                new CreationData().withCreateOption(DiskCreateOption.IMPORT_SECURE).withStorageAccountId(
                    "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount")
                    .withSourceUri("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd")
                    .withSecurityDataUri("https://mystorageaccount.blob.core.windows.net/osimages/vmgs.vhd"))
                .withSecurityProfile(new DiskSecurityProfile().withSecurityType(
                    DiskSecurityTypes.CONFIDENTIAL_VM_VMGUEST_STATE_ONLY_ENCRYPTED_WITH_PLATFORM_KEY)),
            com.azure.core.util.Context.NONE);
    }
}
