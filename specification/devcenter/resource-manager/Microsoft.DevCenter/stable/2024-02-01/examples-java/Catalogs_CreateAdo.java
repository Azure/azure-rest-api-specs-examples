
import com.azure.resourcemanager.devcenter.fluent.models.CatalogInner;
import com.azure.resourcemanager.devcenter.models.CatalogSyncType;
import com.azure.resourcemanager.devcenter.models.GitCatalog;

/**
 * Samples for Catalogs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Catalogs_CreateAdo.json
     */
    /**
     * Sample code: Catalogs_CreateOrUpdateAdo.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsCreateOrUpdateAdo(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().createOrUpdate("rg1", "Contoso", "CentralCatalog",
            new CatalogInner()
                .withAdoGit(new GitCatalog()
                    .withUri("https://contoso@dev.azure.com/contoso/contosoOrg/_git/centralrepo-fakecontoso")
                    .withBranch("main").withSecretIdentifier("fakeTokenPlaceholder").withPath("/templates"))
                .withSyncType(CatalogSyncType.SCHEDULED),
            com.azure.core.util.Context.NONE);
    }
}
