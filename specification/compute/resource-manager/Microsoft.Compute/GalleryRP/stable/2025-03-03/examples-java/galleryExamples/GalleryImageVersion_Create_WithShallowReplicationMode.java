
import com.azure.resourcemanager.compute.fluent.models.GalleryImageVersionInner;
import com.azure.resourcemanager.compute.models.GalleryArtifactVersionFullSource;
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionSafetyProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.ReplicationMode;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;

/**
 * Samples for GalleryImageVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImageVersion_Create_WithShallowReplicationMode.json
     */
    /**
     * Sample code: Create or update a simple Gallery Image Version using shallow replication mode.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGalleryImageVersionUsingShallowReplicationMode(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions()
            .createOrUpdate("myResourceGroup", "myGalleryName", "myGalleryImageName", "1.0.0",
                new GalleryImageVersionInner().withLocation("West US")
                    .withPublishingProfile(new GalleryImageVersionPublishingProfile()
                        .withTargetRegions(Arrays.asList(new TargetRegion().withName("West US")
                            .withRegionalReplicaCount(1).withExcludeFromLatest(false)))
                        .withReplicationMode(ReplicationMode.SHALLOW))
                    .withStorageProfile(new GalleryImageVersionStorageProfile()
                        .withSource(new GalleryArtifactVersionFullSource().withId(
                            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/images/{imageName}")))
                    .withSafetyProfile(new GalleryImageVersionSafetyProfile()
                        .withAllowDeletionOfReplicatedLocations(false).withBlockDeletionBeforeEndOfLife(false)),
                com.azure.core.util.Context.NONE);
    }
}
