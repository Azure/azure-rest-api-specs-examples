
import com.azure.resourcemanager.oracledatabase.models.SystemShapes;

/**
 * Samples for FlexComponents ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/FlexComponents_ListByParent_MaximumSet_Gen.json
     */
    /**
     * Sample code: FlexComponents_ListByParent_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        flexComponentsListByParentMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.flexComponents().listByParent("eastus", SystemShapes.EXADATA_X9M, com.azure.core.util.Context.NONE);
    }
}
