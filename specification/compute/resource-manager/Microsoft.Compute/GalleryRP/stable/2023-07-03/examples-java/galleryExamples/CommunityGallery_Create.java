
import com.azure.resourcemanager.compute.fluent.models.GalleryInner;
import com.azure.resourcemanager.compute.models.CommunityGalleryInfo;
import com.azure.resourcemanager.compute.models.GallerySharingPermissionTypes;
import com.azure.resourcemanager.compute.models.SharingProfile;

/**
 * Samples for Galleries CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2023-07-03/examples/galleryExamples/
     * CommunityGallery_Create.json
     */
    /**
     * Sample code: Create a community gallery.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createACommunityGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleries().createOrUpdate("myResourceGroup",
            "myGalleryName",
            new GalleryInner().withLocation("West US").withDescription("This is the gallery description.")
                .withSharingProfile(new SharingProfile().withPermissions(GallerySharingPermissionTypes.COMMUNITY)
                    .withCommunityGalleryInfo(new CommunityGalleryInfo().withPublisherUri("uri")
                        .withPublisherContact("pir@microsoft.com").withEula("eula").withPublicNamePrefix("PirPublic"))),
            com.azure.core.util.Context.NONE);
    }
}
