Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
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
```
