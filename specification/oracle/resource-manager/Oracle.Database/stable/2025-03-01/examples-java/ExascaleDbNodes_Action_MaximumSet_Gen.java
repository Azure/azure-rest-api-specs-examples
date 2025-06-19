
import com.azure.resourcemanager.oracledatabase.models.DbNodeAction;
import com.azure.resourcemanager.oracledatabase.models.DbNodeActionEnum;

/**
 * Samples for ExascaleDbNodes Action.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/ExascaleDbNodes_Action_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbNodes_Action_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        exascaleDbNodesActionMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exascaleDbNodes().action("rgopenapi", "vmClusterName", "dbNodeName",
            new DbNodeAction().withAction(DbNodeActionEnum.START), com.azure.core.util.Context.NONE);
    }
}
