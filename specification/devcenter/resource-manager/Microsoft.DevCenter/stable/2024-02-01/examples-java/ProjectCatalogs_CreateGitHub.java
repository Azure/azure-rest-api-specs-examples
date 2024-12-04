
import com.azure.resourcemanager.devcenter.models.GitCatalog;

/**
 * Samples for ProjectCatalogs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2024-02-01/examples/
     * ProjectCatalogs_CreateGitHub.json
     */
    /**
     * Sample code: ProjectCatalogs_CreateOrUpdateGitHub.
     * 
     * @param manager Entry point to DevCenterManager.
     */
    public static void
        projectCatalogsCreateOrUpdateGitHub(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.projectCatalogs().define("CentralCatalog").withExistingProject("rg1", "DevProject")
            .withGitHub(new GitCatalog().withUri("https://github.com/Contoso/centralrepo-fake.git").withBranch("main")
                .withSecretIdentifier("fakeTokenPlaceholder").withPath("/templates"))
            .create();
    }
}
