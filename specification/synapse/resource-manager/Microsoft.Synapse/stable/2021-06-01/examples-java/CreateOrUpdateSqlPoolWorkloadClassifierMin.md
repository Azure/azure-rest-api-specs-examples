Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for SqlPoolWorkloadClassifier CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateOrUpdateSqlPoolWorkloadClassifierMin.json
     */
    /**
     * Sample code: Create a workload classifier with the required properties specified.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createAWorkloadClassifierWithTheRequiredPropertiesSpecified(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolWorkloadClassifiers()
            .define("wlm_workloadclassifier")
            .withExistingWorkloadGroup("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", "wlm_workloadgroup")
            .withMemberName("dbo")
            .create();
    }
}
```
