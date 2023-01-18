import com.azure.resourcemanager.datamigration.models.Project;
import com.azure.resourcemanager.datamigration.models.ProjectSourcePlatform;
import com.azure.resourcemanager.datamigration.models.ProjectTargetPlatform;

/** Samples for Projects Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/datamigration/resource-manager/Microsoft.DataMigration/stable/2018-04-19/examples/Projects_Update.json
     */
    /**
     * Sample code: Projects_Update.
     *
     * @param manager Entry point to DataMigrationManager.
     */
    public static void projectsUpdate(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        Project resource =
            manager
                .projects()
                .getWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkProject", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withSourcePlatform(ProjectSourcePlatform.SQL)
            .withTargetPlatform(ProjectTargetPlatform.SQLDB)
            .apply();
    }
}
