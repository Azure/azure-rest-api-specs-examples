import com.azure.resourcemanager.devcenter.models.Catalog;
import com.azure.resourcemanager.devcenter.models.GitCatalog;

/** Samples for Catalogs Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Catalogs_Patch.json
     */
    /**
     * Sample code: Catalogs_Update.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void catalogsUpdate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        Catalog resource =
            manager
                .catalogs()
                .getWithResponse("rg1", "Contoso", "CentralCatalog", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withGitHub(new GitCatalog().withPath("/environments")).apply();
    }
}
