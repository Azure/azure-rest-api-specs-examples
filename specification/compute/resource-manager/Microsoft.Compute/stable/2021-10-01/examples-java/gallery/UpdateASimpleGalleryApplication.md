```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.GalleryApplicationUpdate;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/** Samples for GalleryApplications Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/UpdateASimpleGalleryApplication.json
     */
    /**
     * Sample code: Update a simple gallery Application.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASimpleGalleryApplication(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryApplications()
            .update(
                "myResourceGroup",
                "myGalleryName",
                "myGalleryApplicationName",
                new GalleryApplicationUpdate()
                    .withDescription("This is the gallery application description.")
                    .withEula("This is the gallery application EULA.")
                    .withPrivacyStatementUri("myPrivacyStatementUri}")
                    .withReleaseNoteUri("myReleaseNoteUri")
                    .withSupportedOSType(OperatingSystemTypes.WINDOWS),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
