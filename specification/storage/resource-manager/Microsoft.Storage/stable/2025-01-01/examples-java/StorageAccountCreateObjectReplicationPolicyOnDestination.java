
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
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2025-01-01/examples/
     * StorageAccountCreateObjectReplicationPolicyOnDestination.json
     */
    /**
     * Sample code: StorageAccountCreateObjectReplicationPolicyOnDestination.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        storageAccountCreateObjectReplicationPolicyOnDestination(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getObjectReplicationPoliciesOperations()
            .createOrUpdateWithResponse("res7687", "dst112", "default", new ObjectReplicationPolicyInner()
                .withSourceAccount("src1122").withDestinationAccount("dst112")
                .withRules(Arrays.asList(new ObjectReplicationPolicyRule().withSourceContainer("scont139")
                    .withDestinationContainer("dcont139")
                    .withFilters(new ObjectReplicationPolicyFilter().withPrefixMatch(Arrays.asList("blobA", "blobB")))))
                .withMetrics(new ObjectReplicationPolicyPropertiesMetrics().withEnabled(true)),
                com.azure.core.util.Context.NONE);
    }
}
