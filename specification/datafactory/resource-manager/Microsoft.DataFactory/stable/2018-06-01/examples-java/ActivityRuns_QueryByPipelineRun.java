
import com.azure.resourcemanager.datafactory.models.RunFilterParameters;
import java.time.OffsetDateTime;

/**
 * Samples for ActivityRuns QueryByPipelineRunSync.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * ActivityRuns_QueryByPipelineRun.json
     */
    /**
     * Sample code: ActivityRuns_QueryByPipelineRun.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        activityRunsQueryByPipelineRun(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.activityRuns().queryByPipelineRunWithResponse("exampleResourceGroup", "exampleFactoryName",
            "2f7fdb90-5df1-4b8e-ac2f-064cfa58202b",
            new RunFilterParameters().withLastUpdatedAfter(OffsetDateTime.parse("2018-06-16T00:36:44.3345758Z"))
                .withLastUpdatedBefore(OffsetDateTime.parse("2018-06-16T00:49:48.3686473Z")),
            com.azure.core.util.Context.NONE);
    }
}
