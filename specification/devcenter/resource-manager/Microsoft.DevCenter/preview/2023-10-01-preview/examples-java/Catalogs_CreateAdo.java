import com.azure.resourcemanager.devcenter.models.CatalogSyncType;
import com.azure.resourcemanager.devcenter.models.GitCatalog;

/** Samples for Catalogs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/Catalogs_CreateAdo.json
     */
    /**
     * Sample code: Catalogs_CreateOrUpdateAdo.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsCreateOrUpdateAdo(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .catalogs()
            .define("CentralCatalog")
            .withExistingDevcenter("rg1", "Contoso")
            .withAdoGit(
                new GitCatalog()
                    .withUri("https://contoso@dev.azure.com/contoso/contosoOrg/_git/centralrepo-fakecontoso")
                    .withBranch("main")
                    .withSecretIdentifier("fakeTokenPlaceholder")
                    .withPath("/templates"))
            .withSyncType(CatalogSyncType.SCHEDULED)
            .create();
    }
}
