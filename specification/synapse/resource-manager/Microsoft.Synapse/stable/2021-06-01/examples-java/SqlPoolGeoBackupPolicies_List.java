import com.azure.core.util.Context;

/** Samples for SqlPoolGeoBackupPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolGeoBackupPolicies_List.json
     */
    /**
     * Sample code: Get Sql pool geo backup policy.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getSqlPoolGeoBackupPolicy(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolGeoBackupPolicies().list("sqlcrudtest-4799", "sqlcrudtest-5961", "testdw", Context.NONE);
    }
}
