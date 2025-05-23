
import com.azure.resourcemanager.compute.fluent.models.GalleryImageVersionInner;
import com.azure.resourcemanager.compute.models.DataDiskImageEncryption;
import com.azure.resourcemanager.compute.models.EncryptionImages;
import com.azure.resourcemanager.compute.models.GalleryDataDiskImage;
import com.azure.resourcemanager.compute.models.GalleryDiskImageSource;
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionSafetyProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.GalleryOSDiskImage;
import com.azure.resourcemanager.compute.models.HostCaching;
import com.azure.resourcemanager.compute.models.OSDiskImageEncryption;
import com.azure.resourcemanager.compute.models.StorageAccountType;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;

/**
 * Samples for GalleryImageVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/
     * GalleryImageVersion_Create_WithVHD.json
     */
    /**
     * Sample code: Create or update a simple Gallery Image Version using vhd as a source.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGalleryImageVersionUsingVhdAsASource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().createOrUpdate("myResourceGroup",
            "myGalleryName", "myGalleryImageName", "1.0.0",
            new GalleryImageVersionInner().withLocation("West US").withPublishingProfile(
                new GalleryImageVersionPublishingProfile().withTargetRegions(Arrays.asList(new TargetRegion()
                    .withName("West US").withRegionalReplicaCount(
                        1)
                    .withEncryption(new EncryptionImages().withOsDiskImage(
                        new OSDiskImageEncryption().withDiskEncryptionSetId(
                            "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"))
                        .withDataDiskImages(Arrays.asList(new DataDiskImageEncryption().withDiskEncryptionSetId(
                            "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet")
                            .withLun(1))))
                    .withExcludeFromLatest(false),
                    new TargetRegion().withName("East US").withRegionalReplicaCount(2)
                        .withStorageAccountType(StorageAccountType.STANDARD_ZRS).withExcludeFromLatest(false))))
                .withStorageProfile(new GalleryImageVersionStorageProfile().withOsDiskImage(new GalleryOSDiskImage()
                    .withHostCaching(HostCaching.READ_ONLY)
                    .withSource(new GalleryDiskImageSource().withUri(
                        "https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd")
                        .withStorageAccountId(
                            "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}")))
                    .withDataDiskImages(Arrays.asList(new GalleryDataDiskImage().withHostCaching(HostCaching.NONE)
                        .withSource(new GalleryDiskImageSource().withUri(
                            "https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd")
                            .withStorageAccountId(
                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"))
                        .withLun(1))))
                .withSafetyProfile(new GalleryImageVersionSafetyProfile().withAllowDeletionOfReplicatedLocations(false)
                    .withBlockDeletionBeforeEndOfLife(false)),
            com.azure.core.util.Context.NONE);
    }
}
