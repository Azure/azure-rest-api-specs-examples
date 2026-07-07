
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
     * x-ms-original-file: 2026-04-01/StorageAccountUpdateObjectReplicationPolicyOnDestination.json
     */
    /**
     * Sample code: StorageAccountUpdateObjectReplicationPolicyOnDestination.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void storageAccountUpdateObjectReplicationPolicyOnDestination(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getObjectReplicationPoliciesOperations()
            .createOrUpdateWithResponse("res7687", "dst112", "2a20bb73-5717-4635-985a-5d4cf777438f",
                new ObjectReplicationPolicyInner().withSourceAccount("src1122").withDestinationAccount("dst112")
                    .withRules(
                        Arrays.asList(
                            new ObjectReplicationPolicyRule().withRuleId("d5d18a48-8801-4554-aeaa-74faf65f5ef9")
                                .withSourceContainer("scont139").withDestinationContainer("dcont139")
                                .withFilters(new ObjectReplicationPolicyFilter()
                                    .withPrefixMatch(Arrays.asList("blobA", "blobB"))),
                            new ObjectReplicationPolicyRule().withSourceContainer("scont179")
                                .withDestinationContainer("dcont179")))
                    .withMetrics(new ObjectReplicationPolicyPropertiesMetrics().withEnabled(true))
                    .withPriorityReplication(
                        new ObjectReplicationPolicyPropertiesPriorityReplication().withEnabled(true))
                    .withTagsReplication(new ObjectReplicationPolicyPropertiesTagsReplication().withEnabled(true)),
                com.azure.core.util.Context.NONE);
    }
}
