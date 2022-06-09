```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.GalleryApplicationInner;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/** Samples for GalleryApplications CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-10-01/examples/gallery/CreateOrUpdateASimpleGalleryApplication.json
     */
    /**
     * Sample code: Create or update a simple gallery Application.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateASimpleGalleryApplication(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryApplications()
            .createOrUpdate(
                "myResourceGroup",
                "myGalleryName",
                "myGalleryApplicationName",
                new GalleryApplicationInner()
                    .withLocation("West US")
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
