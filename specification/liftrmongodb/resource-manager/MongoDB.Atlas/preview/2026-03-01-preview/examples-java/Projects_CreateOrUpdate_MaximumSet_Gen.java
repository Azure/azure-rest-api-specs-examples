
import com.azure.resourcemanager.mongodbatlas.models.ProjectProperties;

/**
 * Samples for Projects CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01-preview/Projects_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Projects_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to MongoDBAtlasManager.
     */
    public static void
        projectsCreateOrUpdateMaximumSet(com.azure.resourcemanager.mongodbatlas.MongoDBAtlasManager manager) {
        manager.projects().define("myProject").withExistingOrganization("rgopenapi", "myOrganization")
            .withProperties(new ProjectProperties()).create();
    }
}
