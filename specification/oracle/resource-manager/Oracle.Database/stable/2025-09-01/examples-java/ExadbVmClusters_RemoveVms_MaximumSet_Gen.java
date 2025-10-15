
import com.azure.resourcemanager.oracledatabase.models.DbNodeDetails;
import com.azure.resourcemanager.oracledatabase.models.RemoveVirtualMachineFromExadbVmClusterDetails;
import java.util.Arrays;

/**
 * Samples for ExadbVmClusters RemoveVms.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExadbVmClusters_RemoveVms_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExadbVmClusters_RemoveVms_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        exadbVmClustersRemoveVmsMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exadbVmClusters().removeVms("rgopenapi", "exadbVmClusterName1",
            new RemoveVirtualMachineFromExadbVmClusterDetails()
                .withDbNodes(Arrays.asList(new DbNodeDetails().withDbNodeId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Oracle.Database/exadbVmClusters/vmCluster/dbNodes/dbNodeName"))),
            com.azure.core.util.Context.NONE);
    }
}
