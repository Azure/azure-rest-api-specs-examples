```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.storage.fluent.models.ObjectReplicationPolicyInner;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyFilter;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyRule;
import java.util.Arrays;

/** Samples for ObjectReplicationPoliciesOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/StorageAccountUpdateObjectReplicationPolicyOnDestination.json
     */
    /**
     * Sample code: StorageAccountUpdateObjectReplicationPolicyOnDestination.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void storageAccountUpdateObjectReplicationPolicyOnDestination(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .storageAccounts()
            .manager()
            .serviceClient()
            .getObjectReplicationPoliciesOperations()
            .createOrUpdateWithResponse(
                "res7687",
                "dst112",
                "2a20bb73-5717-4635-985a-5d4cf777438f",
                new ObjectReplicationPolicyInner()
                    .withSourceAccount("src1122")
                    .withDestinationAccount("dst112")
                    .withRules(
                        Arrays
                            .asList(
                                new ObjectReplicationPolicyRule()
                                    .withRuleId("d5d18a48-8801-4554-aeaa-74faf65f5ef9")
                                    .withSourceContainer("scont139")
                                    .withDestinationContainer("dcont139")
                                    .withFilters(
                                        new ObjectReplicationPolicyFilter()
                                            .withPrefixMatch(Arrays.asList("blobA", "blobB"))),
                                new ObjectReplicationPolicyRule()
                                    .withSourceContainer("scont179")
                                    .withDestinationContainer("dcont179"))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.
