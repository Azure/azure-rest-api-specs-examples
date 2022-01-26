Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.synapse.models.DataMaskingState;

/** Samples for DataMaskingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DataMaskingPolicyCreateOrUpdateMax.json
     */
    /**
     * Sample code: Create or update data masking policy max.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateDataMaskingPolicyMax(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .dataMaskingPolicies()
            .define()
            .withExistingSqlPool("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-331")
            .withDataMaskingState(DataMaskingState.ENABLED)
            .withExemptPrincipals("testuser;")
            .create();
    }
}
```
