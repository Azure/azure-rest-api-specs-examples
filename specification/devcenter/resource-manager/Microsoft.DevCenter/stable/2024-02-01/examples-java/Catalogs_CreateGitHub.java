
import com.azure.resourcemanager.devcenter.fluent.models.CatalogInner;
import com.azure.resourcemanager.devcenter.models.CatalogSyncType;
import com.azure.resourcemanager.devcenter.models.GitCatalog;

/**
 * Samples for Catalogs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Catalogs_CreateGitHub.
     * json
     */
    /**
     * Sample code: Catalogs_CreateOrUpdateGitHub.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsCreateOrUpdateGitHub(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().createOrUpdate("rg1", "Contoso", "CentralCatalog",
            new CatalogInner()
                .withGitHub(new GitCatalog().withUri("https://github.com/Contoso/centralrepo-fake.git")
                    .withBranch("main").withSecretIdentifier("fakeTokenPlaceholder").withPath("/templates"))
                .withSyncType(CatalogSyncType.MANUAL),
            com.azure.core.util.Context.NONE);
    }
}
