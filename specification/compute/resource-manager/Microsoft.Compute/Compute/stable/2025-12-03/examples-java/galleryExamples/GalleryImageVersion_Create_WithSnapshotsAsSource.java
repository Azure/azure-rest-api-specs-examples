
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
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryImageVersion_Create_WithSnapshotsAsSource.json
     */
    /**
     * Sample code: Create or update a simple Gallery Image Version using snapshots as a source.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createOrUpdateASimpleGalleryImageVersionUsingSnapshotsAsASource(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().createOrUpdate("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0",
            new GalleryImageVersionInner().withLocation("West US").withPublishingProfile(
                new GalleryImageVersionPublishingProfile().withTargetRegions(Arrays.asList(new TargetRegion()
                    .withName("West US").withRegionalReplicaCount(1)
                    .withEncryption(new EncryptionImages()
                        .withOsDiskImage(new OSDiskImageEncryption().withDiskEncryptionSetId(
                            "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"))
                        .withDataDiskImages(Arrays.asList(new DataDiskImageEncryption().withDiskEncryptionSetId(
                            "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet")
                            .withLun(1))))
                    .withExcludeFromLatest(false),
                    new TargetRegion().withName("East US").withRegionalReplicaCount(2)
                        .withStorageAccountType(StorageAccountType.STANDARD_ZRS)
                        .withEncryption(new EncryptionImages()
                            .withOsDiskImage(new OSDiskImageEncryption().withDiskEncryptionSetId(
                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"))
                            .withDataDiskImages(Arrays.asList(new DataDiskImageEncryption().withDiskEncryptionSetId(
                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet")
                                .withLun(1))))
                        .withExcludeFromLatest(false))))
                .withStorageProfile(new GalleryImageVersionStorageProfile().withOsDiskImage(new GalleryOSDiskImage()
                    .withHostCaching(HostCaching.READ_ONLY)
                    .withSource(new GalleryDiskImageSource().withId(
                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/snapshots/{osSnapshotName}")))
                    .withDataDiskImages(Arrays.asList(new GalleryDataDiskImage().withHostCaching(HostCaching.NONE)
                        .withSource(new GalleryDiskImageSource().withId(
                            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/disks/{dataDiskName}"))
                        .withLun(1))))
                .withSafetyProfile(new GalleryImageVersionSafetyProfile().withAllowDeletionOfReplicatedLocations(false)
                    .withBlockDeletionBeforeEndOfLife(false)),
            com.azure.core.util.Context.NONE);
    }
}
