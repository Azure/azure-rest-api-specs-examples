
import com.azure.resourcemanager.storage.models.ExecutionTarget;
import com.azure.resourcemanager.storage.models.ExecutionTriggerUpdate;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentUpdateExecutionContext;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentUpdateParameters;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentUpdateProperties;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentUpdateReport;
import com.azure.resourcemanager.storage.models.TriggerParametersUpdate;
import com.azure.resourcemanager.storage.models.TriggerType;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for StorageTaskAssignments Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/storageTaskAssignmentsCrud/
     * PatchStorageTaskAssignment.json
     */
    /**
     * Sample code: PatchStorageTaskAssignment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchStorageTaskAssignment(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getStorageTaskAssignments().update("res4228", "sto4445",
            "myassignment1",
            new StorageTaskAssignmentUpdateParameters().withProperties(new StorageTaskAssignmentUpdateProperties()
                .withEnabled(true).withDescription("My Storage task assignment")
                .withExecutionContext(new StorageTaskAssignmentUpdateExecutionContext()
                    .withTarget(new ExecutionTarget().withPrefix(Arrays.asList("prefix1", "prefix2"))
                        .withExcludePrefix(Arrays.asList()))
                    .withTrigger(new ExecutionTriggerUpdate().withType(TriggerType.RUN_ONCE)
                        .withParameters(new TriggerParametersUpdate()
                            .withStartOn(OffsetDateTime.parse("2022-11-15T21:52:47.8145095Z")))))
                .withReport(new StorageTaskAssignmentUpdateReport().withPrefix("container1"))),
            com.azure.core.util.Context.NONE);
    }
}
