
import com.azure.resourcemanager.datamigration.fluent.models.DeleteNodeInner;

/**
 * Samples for SqlMigrationServices DeleteNode.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datamigration/resource-manager/Microsoft.DataMigration/preview/2025-03-15-preview/examples/
     * DeleteIntegrationRuntimeNode.json
     */
    /**
     * Sample code: Delete the integration runtime node.
     * 
     * @param manager Entry point to DataMigrationManager.
     */
    public static void
        deleteTheIntegrationRuntimeNode(com.azure.resourcemanager.datamigration.DataMigrationManager manager) {
        manager.sqlMigrationServices().deleteNodeWithResponse("testrg", "service1",
            new DeleteNodeInner().withNodeName("nodeName").withIntegrationRuntimeName("IRName"),
            com.azure.core.util.Context.NONE);
    }
}
