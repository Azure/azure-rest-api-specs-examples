import com.azure.core.util.Context;

/** Samples for RestorableDroppedSqlPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RestorableDroppedSqlPoolGet.json
     */
    /**
     * Sample code: Get a restorable dropped Sql pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getARestorableDroppedSqlPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .restorableDroppedSqlPools()
            .getWithResponse(
                "restorabledroppeddatabasetest-1257",
                "restorabledroppeddatabasetest-2389",
                "restorabledroppeddatabasetest-7654,131403269876900000",
                Context.NONE);
    }
}
