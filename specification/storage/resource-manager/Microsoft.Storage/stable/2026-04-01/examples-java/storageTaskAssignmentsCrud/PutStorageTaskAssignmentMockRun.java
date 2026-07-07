
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
     * x-ms-original-file: 2026-04-01/storageTaskAssignmentsCrud/PutStorageTaskAssignmentMockRun.json
     */
    /**
     * Sample code: PutStorageTaskAssignmentMockRun.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void putStorageTaskAssignmentMockRun(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageTaskAssignments().create("res4228", "sto4445", "myassignment1",
            new StorageTaskAssignmentInner().withProperties(new StorageTaskAssignmentProperties().withTaskId(
                "/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/myStorageTask")
                .withEnabled(true).withDescription("My Storage task assignment for testing")
                .withExecutionContext(new StorageTaskAssignmentExecutionContext()
                    .withTarget(new ExecutionTarget().withPrefix(Arrays.asList()).withExcludePrefix(Arrays.asList()))
                    .withTrigger(new ExecutionTrigger().withType(TriggerType.MOCK_RUN).withParameters(
                        new TriggerParameters().withStartOn(OffsetDateTime.parse("2023-01-01T00:00:00.1234567Z")))))
                .withReport(new StorageTaskAssignmentReport().withPrefix("reports"))),
            com.azure.core.util.Context.NONE);
    }
}
