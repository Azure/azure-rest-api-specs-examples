
import com.azure.resourcemanager.storage.fluent.models.StorageTaskAssignmentInner;
import com.azure.resourcemanager.storage.models.ExecutionTrigger;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentExecutionContext;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentProperties;
import com.azure.resourcemanager.storage.models.StorageTaskAssignmentReport;
import com.azure.resourcemanager.storage.models.TriggerParameters;
import com.azure.resourcemanager.storage.models.TriggerType;
import java.time.OffsetDateTime;

/**
 * Samples for StorageTaskAssignments Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/storageTaskAssignmentsCrud/PutStorageTaskAssignmentRequiredProperties.json
     */
    /**
     * Sample code: PutStorageTaskAssignmentRequiredProperties.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void
        putStorageTaskAssignmentRequiredProperties(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getStorageTaskAssignments().create("res4228", "sto4445", "myassignment1",
            new StorageTaskAssignmentInner().withProperties(new StorageTaskAssignmentProperties().withTaskId(
                "/subscriptions/1f31ba14-ce16-4281-b9b4-3e78da6e1616/resourceGroups/res4228/providers/Microsoft.StorageActions/storageTasks/mytask1")
                .withEnabled(true).withDescription("My Storage task assignment")
                .withExecutionContext(new StorageTaskAssignmentExecutionContext()
                    .withTrigger(new ExecutionTrigger().withType(TriggerType.RUN_ONCE).withParameters(
                        new TriggerParameters().withStartOn(OffsetDateTime.parse("2022-11-15T21:52:47.8145095Z")))))
                .withReport(new StorageTaskAssignmentReport().withPrefix("container1"))),
            com.azure.core.util.Context.NONE);
    }
}
