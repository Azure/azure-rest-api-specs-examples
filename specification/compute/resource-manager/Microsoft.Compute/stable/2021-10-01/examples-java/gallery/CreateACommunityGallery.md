```java
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.Context;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.compute.fluent.models.GalleryInner;
import com.azure.resourcemanager.compute.models.GallerySharingPermissionTypes;
import com.azure.resourcemanager.compute.models.SharingProfile;
import java.io.IOException;

/** Samples for Galleries CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateACommunityGallery.json
     */
    /**
     * Sample code: Create a community gallery.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createACommunityGallery(com.azure.resourcemanager.AzureResourceManager azure)
        throws IOException {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleries()
            .createOrUpdate(
                "myResourceGroup",
                "myGalleryName",
                new GalleryInner()
                    .withLocation("West US")
                    .withDescription("This is the gallery description.")
                    .withSharingProfile(
                        new SharingProfile()
                            .withPermissions(GallerySharingPermissionTypes.fromString("Community"))
                            .withCommunityGalleryInfo(
                                SerializerFactory
                                    .createDefaultManagementSerializerAdapter()
                                    .deserialize(
                                        "{\"eula\":\"eula\",\"publicNamePrefix\":\"PirPublic\",\"publisherContact\":\"pir@microsoft.com\",\"publisherUri\":\"uri\"}",
                                        Object.class,
                                        SerializerEncoding.JSON))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
