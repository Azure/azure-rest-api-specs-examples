```java
import com.azure.core.util.Context;

/** Samples for SqlPoolWorkloadClassifier Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolWorkloadGroupWorkloadClassifier.json
     */
    /**
     * Sample code: Get a workload classifier for SQL Analytics pool's workload group.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAWorkloadClassifierForSQLAnalyticsPoolSWorkloadGroup(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolWorkloadClassifiers()
            .getWithResponse(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-9187",
                "wlm_workloadgroup",
                "wlm_classifier",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
