import com.azure.resourcemanager.datamigration.models.ProjectSourcePlatform;
import com.azure.resourcemanager.datamigration.models.ProjectTargetPlatform;

/** Samples for Projects CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Projects_CreateOrUpdate.json
     */
    /**
     * Sample code: Projects_CreateOrUpdate.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void projectsCreateOrUpdate(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager
            .projects()
            .define("DmsSdkProject")
            .withRegion("southcentralus")
            .withExistingService("DmsSdkRg", "DmsSdkService")
            .withSourcePlatform(ProjectSourcePlatform.SQL)
            .withTargetPlatform(ProjectTargetPlatform.SQLDB)
            .create();
    }
}
