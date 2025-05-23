
import com.azure.resourcemanager.compute.fluent.models.GalleryApplicationVersionInner;
import com.azure.resourcemanager.compute.models.GalleryApplicationCustomAction;
import com.azure.resourcemanager.compute.models.GalleryApplicationCustomActionParameter;
import com.azure.resourcemanager.compute.models.GalleryApplicationCustomActionParameterType;
import com.azure.resourcemanager.compute.models.GalleryApplicationVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryApplicationVersionSafetyProfile;
import com.azure.resourcemanager.compute.models.StorageAccountType;
import com.azure.resourcemanager.compute.models.TargetRegion;
import com.azure.resourcemanager.compute.models.UserArtifactManage;
import com.azure.resourcemanager.compute.models.UserArtifactSource;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for GalleryApplicationVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/
     * GalleryApplicationVersion_Create.json
     */
    /**
     * Sample code: Create or update a simple gallery Application Version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateASimpleGalleryApplicationVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryApplicationVersions().createOrUpdate(
            "myResourceGroup", "myGalleryName", "myGalleryApplicationName", "1.0.0",
            new GalleryApplicationVersionInner().withLocation("West US")
                .withPublishingProfile(new GalleryApplicationVersionPublishingProfile()
                    .withTargetRegions(Arrays.asList(new TargetRegion().withName("West US").withRegionalReplicaCount(1)
                        .withStorageAccountType(StorageAccountType.STANDARD_LRS).withExcludeFromLatest(false)))
                    .withReplicaCount(1).withEndOfLifeDate(OffsetDateTime.parse("2019-07-01T07:00:00Z"))
                    .withStorageAccountType(StorageAccountType.STANDARD_LRS)
                    .withSource(new UserArtifactSource().withMediaLink(
                        "https://mystorageaccount.blob.core.windows.net/mycontainer/package.zip?{sasKey}"))
                    .withManageActions(new UserArtifactManage()
                        .withInstall(
                            "powershell -command \"Expand-Archive -Path package.zip -DestinationPath C:\\package\"")
                        .withRemove("del C:\\package "))
                    .withCustomActions(Arrays.asList(new GalleryApplicationCustomAction().withName("myCustomAction")
                        .withScript("myCustomActionScript").withDescription("This is the custom action description.")
                        .withParameters(Arrays
                            .asList(new GalleryApplicationCustomActionParameter().withName("myCustomActionParameter")
                                .withRequired(false).withType(GalleryApplicationCustomActionParameterType.STRING)
                                .withDefaultValue("default value of parameter.")
                                .withDescription("This is the description of the parameter"))))))
                .withSafetyProfile(
                    new GalleryApplicationVersionSafetyProfile().withAllowDeletionOfReplicatedLocations(false)),
            com.azure.core.util.Context.NONE);
    }
}
