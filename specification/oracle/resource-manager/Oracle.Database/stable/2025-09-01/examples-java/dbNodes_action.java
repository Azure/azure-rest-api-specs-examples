
import com.azure.resourcemanager.oracledatabase.models.DbNodeAction;
import com.azure.resourcemanager.oracledatabase.models.DbNodeActionEnum;

/**
 * Samples for DbNodes Action.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/dbNodes_action.json
     */
    /**
     * Sample code: DbNodes_Action.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void dbNodesAction(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbNodes().action("rg000", "cluster1", "ocid1....aaaaaa",
            new DbNodeAction().withAction(DbNodeActionEnum.START), com.azure.core.util.Context.NONE);
    }
}
