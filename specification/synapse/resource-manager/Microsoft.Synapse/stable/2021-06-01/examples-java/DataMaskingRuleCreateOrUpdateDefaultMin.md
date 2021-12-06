Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.synapse.models.DataMaskingFunction;

/** Samples for DataMaskingRules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingRuleCreateOrUpdateDefaultMin.json
     */
    /**
     * Sample code: Create/Update data masking rule for default min.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createUpdateDataMaskingRuleForDefaultMin(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .dataMaskingRules()
            .define("rule1")
            .withExistingSqlPool("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331")
            .withSchemaName("dbo")
            .withTableName("Table_1")
            .withColumnName("test1")
            .withMaskingFunction(DataMaskingFunction.DEFAULT)
            .create();
    }
}
```
