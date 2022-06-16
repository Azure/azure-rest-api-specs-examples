import com.azure.core.util.Context;
import com.azure.resourcemanager.batch.models.AutoScaleSettings;
import com.azure.resourcemanager.batch.models.Pool;
import com.azure.resourcemanager.batch.models.ScaleSettings;

/** Samples for Pool Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/PoolUpdate_EnableAutoScale.json
     */
    /**
     * Sample code: UpdatePool - Enable Autoscale.
     *
     * @param manager Entry point to BatchManager.
     */
    public static void updatePoolEnableAutoscale(com.azure.resourcemanager.batch.BatchManager manager) {
        Pool resource =
            manager
                .pools()
                .getWithResponse("default-azurebatch-japaneast", "sampleacct", "testpool", Context.NONE)
                .getValue();
        resource
            .update()
            .withScaleSettings(
                new ScaleSettings().withAutoScale(new AutoScaleSettings().withFormula("$TargetDedicatedNodes=34")))
            .apply();
    }
}
