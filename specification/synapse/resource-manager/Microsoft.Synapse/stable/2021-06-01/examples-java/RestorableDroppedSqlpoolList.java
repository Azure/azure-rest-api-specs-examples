import com.azure.core.util.Context;

/** Samples for RestorableDroppedSqlPools ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RestorableDroppedSqlpoolList.json
     */
    /**
     * Sample code: Get list of restorable dropped Sql pools.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getListOfRestorableDroppedSqlPools(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .restorableDroppedSqlPools()
            .listByWorkspace("restorabledroppeddatabasetest-1349", "restorabledroppeddatabasetest-1840", Context.NONE);
    }
}
