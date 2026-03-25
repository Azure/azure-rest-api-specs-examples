
import com.azure.resourcemanager.containerregistry.fluent.models.ArchiveInner;
import com.azure.resourcemanager.containerregistry.models.ArchivePackageSourceProperties;
import com.azure.resourcemanager.containerregistry.models.PackageSourceType;

/**
 * Samples for Archives Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01-preview/ArchiveCreate.json
     */
    /**
     * Sample code: ArchiveCreate.
     * 
     * @param manager Entry point to ContainerRegistryManager.
     */
    public static void archiveCreate(com.azure.resourcemanager.containerregistry.ContainerRegistryManager manager) {
        manager.serviceClient().getArchives().create("myResourceGroup", "myRegistry", "rpm", "myArchiveName",
            new ArchiveInner()
                .withPackageSource(
                    new ArchivePackageSourceProperties().withType(PackageSourceType.REMOTE).withUrl("string"))
                .withPublishedVersion("string").withRepositoryEndpointPrefix("string"),
            com.azure.core.util.Context.NONE);
    }
}
