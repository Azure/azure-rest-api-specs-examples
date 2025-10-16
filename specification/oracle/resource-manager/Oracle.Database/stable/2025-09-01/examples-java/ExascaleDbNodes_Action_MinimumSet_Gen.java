
import com.azure.resourcemanager.oracledatabase.models.DbNodeAction;
import com.azure.resourcemanager.oracledatabase.models.DbNodeActionEnum;

/**
 * Samples for ExascaleDbNodes Action.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExascaleDbNodes_Action_MinimumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbNodes_Action_MinimumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        exascaleDbNodesActionMinimumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exascaleDbNodes().action("rgopenapi", "exadbvmcluster1", "exascaledbnode1",
            new DbNodeAction().withAction(DbNodeActionEnum.START), com.azure.core.util.Context.NONE);
    }
}
