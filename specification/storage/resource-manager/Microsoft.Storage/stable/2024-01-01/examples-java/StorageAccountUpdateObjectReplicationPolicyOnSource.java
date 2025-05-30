
import com.azure.resourcemanager.storage.fluent.models.ObjectReplicationPolicyInner;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyFilter;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyPropertiesMetrics;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyRule;
import java.util.Arrays;

/**
 * Samples for ObjectReplicationPoliciesOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/
     * StorageAccountUpdateObjectReplicationPolicyOnSource.json
     */
    /**
     * Sample code: StorageAccountUpdateObjectReplicationPolicyOnSource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        storageAccountUpdateObjectReplicationPolicyOnSource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getObjectReplicationPoliciesOperations()
            .createOrUpdateWithResponse("res7687", "src1122", "2a20bb73-5717-4635-985a-5d4cf777438f",
                new ObjectReplicationPolicyInner().withSourceAccount("src1122").withDestinationAccount("dst112")
                    .withRules(
                        Arrays.asList(
                            new ObjectReplicationPolicyRule().withRuleId("d5d18a48-8801-4554-aeaa-74faf65f5ef9")
                                .withSourceContainer("scont139").withDestinationContainer("dcont139")
                                .withFilters(new ObjectReplicationPolicyFilter()
                                    .withPrefixMatch(Arrays.asList("blobA", "blobB"))),
                            new ObjectReplicationPolicyRule().withRuleId("cfbb4bc2-8b60-429f-b05a-d1e0942b33b2")
                                .withSourceContainer("scont179").withDestinationContainer("dcont179")))
                    .withMetrics(new ObjectReplicationPolicyPropertiesMetrics().withEnabled(true)),
                com.azure.core.util.Context.NONE);
    }
}
