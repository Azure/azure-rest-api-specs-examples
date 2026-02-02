
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionUpdate;
import com.azure.resourcemanager.compute.models.StorageAccountType;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;

/**
 * Samples for GalleryImageVersions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImageVersion_Update_WithoutSourceId.json
     */
    /**
     * Sample code: Update a simple Gallery Image Version without source id.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        updateASimpleGalleryImageVersionWithoutSourceId(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions()
            .update("myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", new GalleryImageVersionUpdate()
                .withPublishingProfile(new GalleryImageVersionPublishingProfile()
                    .withTargetRegions(Arrays.asList(new TargetRegion().withName("West US").withRegionalReplicaCount(1),
                        new TargetRegion().withName("East US").withRegionalReplicaCount(2)
                            .withStorageAccountType(StorageAccountType.STANDARD_ZRS))))
                .withStorageProfile(new GalleryImageVersionStorageProfile()), com.azure.core.util.Context.NONE);
    }
}
