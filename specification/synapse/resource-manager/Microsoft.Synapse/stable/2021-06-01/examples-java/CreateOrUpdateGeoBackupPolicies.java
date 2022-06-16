import com.azure.resourcemanager.synapse.models.GeoBackupPolicyName;
import com.azure.resourcemanager.synapse.models.GeoBackupPolicyState;

/** Samples for SqlPoolGeoBackupPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateGeoBackupPolicies.json
     */
    /**
     * Sample code: Create geo backup policy.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createGeoBackupPolicy(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolGeoBackupPolicies()
            .define(GeoBackupPolicyName.DEFAULT)
            .withExistingSqlPool("testrg", "testws", "testdw")
            .withState(GeoBackupPolicyState.ENABLED)
            .create();
    }
}
