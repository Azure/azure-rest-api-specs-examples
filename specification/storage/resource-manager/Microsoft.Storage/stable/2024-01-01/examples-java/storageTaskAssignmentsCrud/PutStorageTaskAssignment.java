
import com.azure.resourcemanager.storage.fluent.models.StorageTaskAssignmentInner;
import com.azure.resourcemanager.storage.models.ExecutionTarget;
import com.azure.resourcemanager.storage.models.ExecutionTrigger;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentExecutionContext;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentProperties;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentReport;
import com.azure.resourcemanager.storage.models.TriggerParameters;
import com.azure.resourcemanager.storage.models.TriggerType;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for StorageTaskAssignments Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/storageTaskAssignmentsCrud/
     * PutStorageTaskAssignment.json
     */
    /**
     * Sample code: PutStorageTaskAssignment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void putStorageTaskAssignment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageTaskAssignments().create("res4228", "sto4445",
            "myassignment1",
            new StorageTaskAssignmentInner().withProperties(new StorageTaskAssignmentProperties().withTaskId(
                "/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1")
                .withEnabled(true).withDescription("My Storage task assignment")
                .withExecutionContext(new StorageTaskAssignmentExecutionContext()
                    .withTarget(new ExecutionTarget().withPrefix(Arrays.asList("prefix1", "prefix2"))
                        .withExcludePrefix(Arrays.asList()))
                    .withTrigger(new ExecutionTrigger().withType(TriggerType.RUN_ONCE).withParameters(
                        new TriggerParameters().withStartOn(OffsetDateTime.parse("2022-11-15T21:52:47.8145095Z")))))
                .withReport(new StorageTaskAssignmentReport().withPrefix("container1"))),
            com.azure.core.util.Context.NONE);
    }
}
