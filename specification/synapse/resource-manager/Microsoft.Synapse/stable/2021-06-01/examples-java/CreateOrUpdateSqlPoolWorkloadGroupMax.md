```java
/** Samples for SqlPoolWorkloadGroup CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolWorkloadGroupMax.json
     */
    /**
     * Sample code: Create a workload group with all properties specified.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createAWorkloadGroupWithAllPropertiesSpecified(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolWorkloadGroups()
            .define("smallrc")
            .withExistingSqlPool("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187")
            .withMinResourcePercent(0)
            .withMaxResourcePercent(100)
            .withMinResourcePercentPerRequest(3.0)
            .withMaxResourcePercentPerRequest(3.0)
            .withImportance("normal")
            .withQueryExecutionTimeout(0)
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
