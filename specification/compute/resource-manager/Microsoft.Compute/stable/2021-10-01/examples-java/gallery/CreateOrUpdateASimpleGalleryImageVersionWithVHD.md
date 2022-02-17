Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.12.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.GalleryImageVersionInner;
import com.azure.resourcemanager.compute.models.DataDiskImageEncryption;
import com.azure.resourcemanager.compute.models.EncryptionImages;
import com.azure.resourcemanager.compute.models.GalleryArtifactVersionSource;
import com.azure.resourcemanager.compute.models.GalleryDataDiskImage;
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.GalleryOSDiskImage;
import com.azure.resourcemanager.compute.models.HostCaching;
import com.azure.resourcemanager.compute.models.OSDiskImageEncryption;
import com.azure.resourcemanager.compute.models.StorageAccountType;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;

/** Samples for GalleryImageVersions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryImageVersionWithVHD.json
     */
    /**
     * Sample code: Create or update a simple Gallery Image Version using vhd as a source.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGalleryImageVersionUsingVhdAsASource(
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
                                                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myDiskEncryptionSet"))
                                                    .withDataDiskImages(
                                                        Arrays
                                                            .asList(
                                                                new DataDiskImageEncryption()
                                                                    .withDiskEncryptionSetId(
                                                                        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSet/myOtherDiskEncryptionSet")
                                                                    .withLun(1)))),
                                        new TargetRegion()
                                            .withName("East US")
                                            .withRegionalReplicaCount(2)
                                            .withStorageAccountType(StorageAccountType.STANDARD_ZRS))))
                    .withStorageProfile(
                        new GalleryImageVersionStorageProfile()
                            .withOsDiskImage(
                                new GalleryOSDiskImage()
                                    .withHostCaching(HostCaching.READ_ONLY)
                                    .withSource(
                                        new GalleryArtifactVersionSource()
                                            .withId(
                                                "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}")
                                            .withUri(
                                                "https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd")))
                            .withDataDiskImages(
                                Arrays
                                    .asList(
                                        new GalleryDataDiskImage()
                                            .withHostCaching(HostCaching.NONE)
                                            .withSource(
                                                new GalleryArtifactVersionSource()
                                                    .withId(
                                                        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}")
                                                    .withUri(
                                                        "https://gallerysourcencus.blob.core.windows.net/myvhds/Windows-Server-2012-R2-20171216-en.us-128GB.vhd"))
                                            .withLun(1)))),
                Context.NONE);
    }
}
```
