
import com.azure.resourcemanager.compute.fluent.models.GalleryImageVersionInner;
import com.azure.resourcemanager.compute.models.ConfidentialVMEncryptionType;
import com.azure.resourcemanager.compute.models.DataDiskImageEncryption;
import com.azure.resourcemanager.compute.models.DataDiskImageSecurityProfile;
import com.azure.resourcemanager.compute.models.EncryptionImages;
import com.azure.resourcemanager.compute.models.GalleryArtifactVersionFullSource;
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.OSDiskImageEncryption;
import com.azure.resourcemanager.compute.models.OSDiskImageSecurityProfile;
import com.azure.resourcemanager.compute.models.ReplicationMode;
import com.azure.resourcemanager.compute.models.StorageAccountType;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;

/**
 * Samples for GalleryImageVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-03/galleryExamples/GalleryImageVersion_Create_WithCVMDataDiskEncryption.json
     */
    /**
     * Sample code: Create or update a Gallery Image Version with CVM Data Disk Encryption using customer-managed key.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createOrUpdateAGalleryImageVersionWithCVMDataDiskEncryptionUsingCustomerManagedKey(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().createOrUpdate("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0",
            new GalleryImageVersionInner().withLocation("eastus")
                .withPublishingProfile(new GalleryImageVersionPublishingProfile()
                    .withTargetRegions(Arrays.asList(new TargetRegion().withName("eastus").withRegionalReplicaCount(1)
                        .withStorageAccountType(StorageAccountType.STANDARD_ZRS)
                        .withEncryption(new EncryptionImages().withOsDiskImage(
                            new OSDiskImageEncryption().withSecurityProfile(new OSDiskImageSecurityProfile()
                                .withConfidentialVMEncryptionType(ConfidentialVMEncryptionType.ENCRYPTED_WITH_PMK)))
                            .withDataDiskImages(Arrays.asList(new DataDiskImageEncryption()
                                .withSecurityProfile(new DataDiskImageSecurityProfile()
                                    .withConfidentialVMEncryptionType(
                                        ConfidentialVMEncryptionType.DATA_DISK_ENCRYPTED_WITH_CMK)
                                    .withSecureVMDiskEncryptionSetId(
                                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/myDiskEncryptionSet"))
                                .withLun(0))))
                        .withExcludeFromLatest(false)))
                    .withReplicaCount(1).withExcludeFromLatest(false).withReplicationMode(ReplicationMode.FULL))
                .withStorageProfile(new GalleryImageVersionStorageProfile()
                    .withSource(new GalleryArtifactVersionFullSource().withVirtualMachineId(
                        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/virtualMachines/myVM"))),
            com.azure.core.util.Context.NONE);
    }
}
