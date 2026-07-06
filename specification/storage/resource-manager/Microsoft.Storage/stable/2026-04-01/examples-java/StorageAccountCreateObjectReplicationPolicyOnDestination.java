
import com.azure.resourcemanager.storage.fluent.models.ObjectReplicationPolicyInner;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyFilter;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyPropertiesMetrics;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyPropertiesPriorityReplication;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyPropertiesTagsReplication;
import com.azure.resourcemanager.storage.models.ObjectReplicationPolicyRule;
import java.util.Arrays;

/**
 * Samples for ObjectReplicationPoliciesOperation CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/StorageAccountCreateObjectReplicationPolicyOnDestination.json
     */
    /**
     * Sample code: StorageAccountCreateObjectReplicationPolicyOnDestination.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountCreateObjectReplicationPolicyOnDestination(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getObjectReplicationPoliciesOperations().createOrUpdateWithResponse("res7687", "dst112",
            "default",
            new ObjectReplicationPolicyInner().withSourceAccount("src1122").withDestinationAccount("dst112")
                .withRules(Arrays.asList(new ObjectReplicationPolicyRule().withSourceContainer("scont139")
                    .withDestinationContainer("dcont139")
                    .withFilters(new ObjectReplicationPolicyFilter().withPrefixMatch(Arrays.asList("blobA", "blobB")))))
                .withMetrics(new ObjectReplicationPolicyPropertiesMetrics().withEnabled(true))
                .withPriorityReplication(new ObjectReplicationPolicyPropertiesPriorityReplication().withEnabled(true))
                .withTagsReplication(new ObjectReplicationPolicyPropertiesTagsReplication().withEnabled(true)),
            com.azure.core.util.Context.NONE);
    }
}
