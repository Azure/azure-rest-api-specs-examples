
import com.azure.resourcemanager.devcenter.models.GitCatalog;

/**
 * Samples for ProjectCatalogs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/ProjectCatalogs_CreateAdo
     * .json
     */
    /**
     * Sample code: ProjectCatalogs_CreateOrUpdateAdo.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void projectCatalogsCreateOrUpdateAdo(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectCatalogs().define("CentralCatalog").withExistingProject("rg1", "DevProject")
            .withAdoGit(new GitCatalog()
                .withUri("https://contoso@dev.azure.com/contoso/contosoOrg/_git/centralrepo-fakecontoso")
                .withBranch("main").withSecretIdentifier("fakeTokenPlaceholder").withPath("/templates"))
            .create();
    }
}
