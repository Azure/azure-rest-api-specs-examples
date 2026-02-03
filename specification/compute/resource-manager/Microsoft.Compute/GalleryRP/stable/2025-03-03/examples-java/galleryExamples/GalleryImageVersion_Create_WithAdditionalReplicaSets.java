
import com.azure.resourcemanager.compute.fluent.models.GalleryImageVersionInner;
import com.azure.resourcemanager.compute.models.AdditionalReplicaSet;
import com.azure.resourcemanager.compute.models.DataDiskImageEncryption;
import com.azure.resourcemanager.compute.models.EncryptionImages;
import com.azure.resourcemanager.compute.models.GalleryArtifactVersionFullSource;
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionSafetyProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
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
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImageVersion_Create_WithAdditionalReplicaSets.json
     */
    /**
     * Sample code: Create or update a simple Gallery Image Version with Direct Drive replicas.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGalleryImageVersionWithDirectDriveReplicas(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().createOrUpdate("myResourceGroup",
            "myGalleryName", "myGalleryImageName", "1.0.0",
            new GalleryImageVersionInner().withLocation("West US").withPublishingProfile(
                new GalleryImageVersionPublishingProfile().withTargetRegions(Arrays.asList(new TargetRegion()
                    .withName("West US").withRegionalReplicaCount(
                        1)
                    .withEncryption(new EncryptionImages().withOsDiskImage(
                        new OSDiskImageEncryption().withDiskEncryptionSetId(
                            "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"))
                        .withDataDiskImages(Arrays.asList(new DataDiskImageEncryption().withDiskEncryptionSetId(
                            "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet")
                            .withLun(0),
                            new DataDiskImageEncryption().withDiskEncryptionSetId(
                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet")
                                .withLun(1))))
                    .withExcludeFromLatest(false)
                    .withAdditionalReplicaSets(Arrays.asList(new AdditionalReplicaSet()
                        .withStorageAccountType(StorageAccountType.fromString("PreviumV2_LRS"))
                        .withRegionalReplicaCount(1))),
                    new TargetRegion().withName("East US").withRegionalReplicaCount(2)
                        .withStorageAccountType(StorageAccountType.STANDARD_ZRS)
                        .withEncryption(new EncryptionImages()
                            .withOsDiskImage(new OSDiskImageEncryption().withDiskEncryptionSetId(
                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"))
                            .withDataDiskImages(Arrays.asList(new DataDiskImageEncryption().withDiskEncryptionSetId(
                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet")
                                .withLun(0),
                                new DataDiskImageEncryption().withDiskEncryptionSetId(
                                    "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet")
                                    .withLun(1))))
                        .withExcludeFromLatest(false))))
                .withStorageProfile(
                    new GalleryImageVersionStorageProfile().withSource(new GalleryArtifactVersionFullSource().withId(
                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}")))
                .withSafetyProfile(
                    new GalleryImageVersionSafetyProfile().withAllowDeletionOfReplicatedLocations(false)),
            com.azure.core.util.Context.NONE);
    }
}
