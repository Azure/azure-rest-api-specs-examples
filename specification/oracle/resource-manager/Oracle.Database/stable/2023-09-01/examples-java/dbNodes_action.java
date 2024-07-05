
import com.azure.resourcemanager.oracledatabase.models.DbNodeAction;
import com.azure.resourcemanager.oracledatabase.models.DbNodeActionEnum;

/**
 * Samples for DbNodes Action.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dbNodes_action.json
     */
    /**
     * Sample code: VM actions on DbNodes of VM Cluster.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        vmActionsOnDbNodesOfVMCluster(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbNodes().action("rg000", "cluster1", "ocid1....aaaaaa",
            new DbNodeAction().withAction(DbNodeActionEnum.START), com.azure.core.util.Context.NONE);
    }
}
