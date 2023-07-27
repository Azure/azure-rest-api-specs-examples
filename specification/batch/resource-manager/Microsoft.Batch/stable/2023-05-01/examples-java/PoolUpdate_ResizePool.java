import com.azure.resourcemanager.batch.models.ComputeNodeDeallocationOption;
import com.azure.resourcemanager.batch.models.FixedScaleSettings;
import com.azure.resourcemanager.batch.models.Pool;
import com.azure.resourcemanager.batch.models.ScaleSettings;
import java.time.Duration;

/** Samples for Pool Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/PoolUpdate_ResizePool.json
     */
    /**
     * Sample code: UpdatePool - Resize Pool.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void updatePoolResizePool(com.azure.resourcemanager.batch.BatchManager manager) {
        Pool resource =
            manager
                .pools()
                .getWithResponse(
                    "default-azurebatch-japaneast", "sampleacct", "testpool", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withScaleSettings(
                new ScaleSettings()
                    .withFixedScale(
                        new FixedScaleSettings()
                            .withResizeTimeout(Duration.parse("PT8M"))
                            .withTargetDedicatedNodes(5)
                            .withTargetLowPriorityNodes(0)
                            .withNodeDeallocationOption(ComputeNodeDeallocationOption.TASK_COMPLETION)))
            .apply();
    }
}
