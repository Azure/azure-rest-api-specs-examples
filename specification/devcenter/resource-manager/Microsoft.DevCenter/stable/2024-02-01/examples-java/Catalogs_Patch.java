
import com.azure.resourcemanager.devcenter.models.CatalogSyncType;
import com.azure.resourcemanager.devcenter.models.CatalogUpdate;
import com.azure.resourcemanager.devcenter.models.GitCatalog;

/**
 * Samples for Catalogs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/Catalogs_Patch.json
     */
    /**
     * Sample code: Catalogs_Update.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsUpdate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.catalogs().update(
            "rg1", "Contoso", "CentralCatalog", new CatalogUpdate()
                .withGitHub(new GitCatalog().withPath("/environments")).withSyncType(CatalogSyncType.SCHEDULED),
            com.azure.core.util.Context.NONE);
    }
}
