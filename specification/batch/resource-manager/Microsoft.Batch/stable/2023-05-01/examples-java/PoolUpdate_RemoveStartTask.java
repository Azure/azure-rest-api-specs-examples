import com.azure.resourcemanager.batch.models.Pool;
import com.azure.resourcemanager.batch.models.StartTask;

/** Samples for Pool Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolUpdate_RemoveStartTask.json
     */
    /**
     * Sample code: UpdatePool - Remove Start Task.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void updatePoolRemoveStartTask(com.azure.resourcemanager.batch.BatchManager manager) {
        Pool resource =
            manager
                .pools()
                .getWithResponse(
                    "default-azurebatch-japaneast", "sampleacct", "testpool", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withStartTask(new StartTask()).apply();
    }
}
