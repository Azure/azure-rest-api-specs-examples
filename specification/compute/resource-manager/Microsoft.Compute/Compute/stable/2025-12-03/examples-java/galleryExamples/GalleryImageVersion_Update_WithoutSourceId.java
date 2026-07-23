
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
     * x-ms-original-file: 2025-12-03/galleryExamples/GalleryImageVersion_Update_WithoutSourceId.json
     */
    /**
     * Sample code: Update a simple Gallery Image Version without source id.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        updateASimpleGalleryImageVersionWithoutSourceId(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions()
            .update("myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0", new GalleryImageVersionUpdate()
                .withPublishingProfile(new GalleryImageVersionPublishingProfile()
                    .withTargetRegions(Arrays.asList(new TargetRegion().withName("West US").withRegionalReplicaCount(1),
                        new TargetRegion().withName("East US").withRegionalReplicaCount(2)
                            .withStorageAccountType(StorageAccountType.STANDARD_ZRS))))
                .withStorageProfile(new GalleryImageVersionStorageProfile()), com.azure.core.util.Context.NONE);
    }
}
