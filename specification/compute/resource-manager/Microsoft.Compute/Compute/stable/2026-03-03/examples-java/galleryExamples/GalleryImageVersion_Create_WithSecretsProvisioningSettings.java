
import com.azure.resourcemanager.compute.fluent.models.GalleryImageVersionInner;
import com.azure.resourcemanager.compute.models.GalleryDiskImageSource;
import com.azure.resourcemanager.compute.models.GalleryImageVersionPublishingProfile;
import com.azure.resourcemanager.compute.models.GalleryImageVersionStorageProfile;
import com.azure.resourcemanager.compute.models.GalleryOSDiskImage;
import com.azure.resourcemanager.compute.models.HostCaching;
import com.azure.resourcemanager.compute.models.ImageVersionSecurityProfile;
import com.azure.resourcemanager.compute.models.SecretsProvisioningComponent;
import com.azure.resourcemanager.compute.models.SecretsProvisioningComponentName;
import com.azure.resourcemanager.compute.models.SecretsProvisioningSettings;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;

/**
 * Samples for GalleryImageVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-03/galleryExamples/GalleryImageVersion_Create_WithSecretsProvisioningSettings.json
     */
    /**
     * Sample code: Create or update a simple Gallery Image Version with secrets provisioning settings.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createOrUpdateASimpleGalleryImageVersionWithSecretsProvisioningSettings(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryImageVersions().createOrUpdate("myResourceGroup", "myGalleryName",
            "myGalleryImageName", "1.0.0",
            new GalleryImageVersionInner().withLocation("West US")
                .withPublishingProfile(new GalleryImageVersionPublishingProfile().withTargetRegions(Arrays.asList(
                    new TargetRegion().withName("West US").withRegionalReplicaCount(1).withExcludeFromLatest(false))))
                .withStorageProfile(new GalleryImageVersionStorageProfile().withOsDiskImage(new GalleryOSDiskImage()
                    .withHostCaching(HostCaching.READ_ONLY)
                    .withSource(new GalleryDiskImageSource()
                        .withUri("https://gallerysourcencus.blob.core.windows.net/myvhds/Linux-VM-2024.vhd")
                        .withStorageAccountId(
                            "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/{storageAccount}"))))
                .withSecurityProfile(new ImageVersionSecurityProfile().withSecretsProvisioningSettings(
                    new SecretsProvisioningSettings().withIsSupported(true).withOsName("mariner")
                        .withComponents(Arrays.asList(
                            new SecretsProvisioningComponent()
                                .withName(SecretsProvisioningComponentName.AZURE_GUEST_AGENT).withVersion("2.7.0"),
                            new SecretsProvisioningComponent()
                                .withName(SecretsProvisioningComponentName.SECRETS_PROVISIONING_LIBRARY)
                                .withVersion("1.0.0"))))),
            com.azure.core.util.Context.NONE);
    }
}
