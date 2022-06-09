```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.GeoBackupPolicyName;

/** Samples for SqlPoolGeoBackupPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolGeoBackupPolicy.json
     */
    /**
     * Sample code: Get Sql pool geo backup policy.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getSqlPoolGeoBackupPolicy(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolGeoBackupPolicies()
            .getWithResponse(
                "sqlcrudtest-4799", "sqlcrudtest-5961", "testdw", GeoBackupPolicyName.DEFAULT, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
