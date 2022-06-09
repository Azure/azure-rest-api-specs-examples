```java
import com.azure.core.util.Context;

/** Samples for SqlPoolWorkloadClassifier Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPoolWorkloadGroupWorkloadClassifer.json
     */
    /**
     * Sample code: Delete a workload classifier of a SQL Analytics pool's workload group.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteAWorkloadClassifierOfASQLAnalyticsPoolSWorkloadGroup(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolWorkloadClassifiers()
            .delete(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-9187",
                "wlm_workloadgroup",
                "wlm_workloadclassifier",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
