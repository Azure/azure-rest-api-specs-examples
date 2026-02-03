
import com.azure.resourcemanager.compute.fluent.models.GalleryImageVersionInner;
import com.azure.resourcemanager.compute.models.GalleryArtifactVersionFullSource;
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.StorageAccountStrategy;
import com.azure.resourcemanager.compute.models.StorageAccountType;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;

/**
 * Samples for GalleryImageVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/galleryExamples/
     * GalleryImageVersion_Create_WithStorageAccountStrategy.json
     */
    /**
     * Sample code: Create or update a simple Gallery Image Version with StorageAccountStrategy and regional
     * StorageAccountType override.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateASimpleGalleryImageVersionWithStorageAccountStrategyAndRegionalStorageAccountTypeOverride(
            com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryImageVersions().createOrUpdate("myResourceGroup",
            "myGalleryName", "myGalleryImageName", "1.0.0",
            new GalleryImageVersionInner().withLocation("West US")
                .withPublishingProfile(new GalleryImageVersionPublishingProfile()
                    .withTargetRegions(
                        Arrays.asList(new TargetRegion().withName("West US"), new TargetRegion().withName("East US"),
                            new TargetRegion().withName("East US 2")
                                .withStorageAccountType(StorageAccountType.PREMIUM_LRS)))
                    .withStorageAccountStrategy(StorageAccountStrategy.PREFER_STANDARD_ZRS))
                .withStorageProfile(new GalleryImageVersionStorageProfile()
                    .withSource(new GalleryArtifactVersionFullSource().withVirtualMachineId(
                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Compute/virtualMachines/{vmName}"))),
            com.azure.core.util.Context.NONE);
    }
}
