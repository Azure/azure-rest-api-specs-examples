Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.SharingUpdateInner;
import com.azure.resourcemanager.compute.models.SharingProfileGroup;
import com.azure.resourcemanager.compute.models.SharingProfileGroupTypes;
import com.azure.resourcemanager.compute.models.SharingUpdateOperationTypes;
import java.util.Arrays;

/** Samples for GallerySharingProfile Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/gallery/AddToSharingProfileInAGallery.json
     */
    /**
     * Sample code: Add sharing id to the sharing profile of a gallery.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void addSharingIdToTheSharingProfileOfAGallery(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGallerySharingProfiles()
            .update(
                "myResourceGroup",
                "myGalleryName",
                new SharingUpdateInner()
                    .withOperationType(SharingUpdateOperationTypes.ADD)
                    .withGroups(
                        Arrays
                            .asList(
                                new SharingProfileGroup()
                                    .withType(SharingProfileGroupTypes.SUBSCRIPTIONS)
                                    .withIds(
                                        Arrays
                                            .asList(
                                                "34a4ab42-0d72-47d9-bd1a-aed207386dac",
                                                "380fd389-260b-41aa-bad9-0a83108c370b")),
                                new SharingProfileGroup()
                                    .withType(SharingProfileGroupTypes.AADTENANTS)
                                    .withIds(Arrays.asList("c24c76aa-8897-4027-9b03-8f7928b54ff6")))),
                Context.NONE);
    }
}
```
