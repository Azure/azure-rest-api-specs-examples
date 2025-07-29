
import com.azure.resourcemanager.datamigration.models.ProjectFile;
import com.azure.resourcemanager.datamigration.models.ProjectFileProperties;

/**
 * Samples for Files Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * Files_Update.json
     */
    /**
     * Sample code: Files_Update.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void filesUpdate(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        ProjectFile resource = manager.files().getWithResponse("DmsSdkRg", "DmsSdkService", "DmsSdkProject",
            "x114d023d8", com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new ProjectFileProperties().withFilePath("DmsSdkFilePath/DmsSdkFile.sql"))
            .apply();
    }
}
