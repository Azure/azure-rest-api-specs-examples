import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.GalleryImageVersionInner;
import com.azure.resourcemanager.compute.models.DataDiskImageEncryption;
import com.azure.resourcemanager.compute.models.EncryptionImages;
import com.azure.resourcemanager.compute.models.GalleryArtifactVersionSource;
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.OSDiskImageEncryption;
import com.azure.resourcemanager.compute.models.StorageAccountType;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;

/** Samples for GalleryImageVersions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryImageVersionWithTargetExtendedLocations.json
     */
    /**
     * Sample code: Create or update a simple gallery image version with target extended locations specified.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGalleryImageVersionWithTargetExtendedLocationsSpecified(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryImageVersions()
            .createOrUpdate(
                "myResourceGroup",
                "myGalleryName",
                "myGalleryImageName",
                "1.0.0",
                new GalleryImageVersionInner()
                    .withLocation("West US")
                    .withPublishingProfile(
                        new GalleryImageVersionPublishingProfile()
                            .withTargetRegions(
                                Arrays
                                    .asList(
                                        new TargetRegion()
                                            .withName("West US")
                                            .withRegionalReplicaCount(1)
                                            .withEncryption(
                                                new EncryptionImages()
                                                    .withOsDiskImage(
                                                        new OSDiskImageEncryption()
                                                            .withDiskEncryptionSetId(
                                                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet"))
                                                    .withDataDiskImages(
                                                        Arrays
                                                            .asList(
                                                                new DataDiskImageEncryption()
                                                                    .withDiskEncryptionSetId(
                                                                        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherWestUSDiskEncryptionSet")
                                                                    .withLun(0),
                                                                new DataDiskImageEncryption()
                                                                    .withDiskEncryptionSetId(
                                                                        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myWestUSDiskEncryptionSet")
                                                                    .withLun(1)))),
                                        new TargetRegion()
                                            .withName("East US")
                                            .withRegionalReplicaCount(2)
                                            .withStorageAccountType(StorageAccountType.STANDARD_ZRS)
                                            .withEncryption(
                                                new EncryptionImages()
                                                    .withOsDiskImage(
                                                        new OSDiskImageEncryption()
                                                            .withDiskEncryptionSetId(
                                                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet"))
                                                    .withDataDiskImages(
                                                        Arrays
                                                            .asList(
                                                                new DataDiskImageEncryption()
                                                                    .withDiskEncryptionSetId(
                                                                        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherEastUSDiskEncryptionSet")
                                                                    .withLun(0),
                                                                new DataDiskImageEncryption()
                                                                    .withDiskEncryptionSetId(
                                                                        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myEastUSDiskEncryptionSet")
                                                                    .withLun(1)))))))
                    .withStorageProfile(
                        new GalleryImageVersionStorageProfile()
                            .withSource(
                                new GalleryArtifactVersionSource()
                                    .withId(
                                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}"))),
                Context.NONE);
    }
}
