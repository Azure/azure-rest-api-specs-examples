```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.datafactory.models.RunFilterParameters;
import com.azure.resourcemanager.datafactory.models.RunQueryFilter;
import com.azure.resourcemanager.datafactory.models.RunQueryFilterOperand;
import com.azure.resourcemanager.datafactory.models.RunQueryFilterOperator;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for PipelineRuns QueryByFactory. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_QueryByFactory.json
     */
    /**
     * Sample code: PipelineRuns_QueryByFactory.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void pipelineRunsQueryByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .pipelineRuns()
            .queryByFactoryWithResponse(
                "exampleResourceGroup",
                "exampleFactoryName",
                new RunFilterParameters()
                    .withLastUpdatedAfter(OffsetDateTime.parse("2018-06-16T00:36:44.3345758Z"))
                    .withLastUpdatedBefore(OffsetDateTime.parse("2018-06-16T00:49:48.3686473Z"))
                    .withFilters(
                        Arrays
                            .asList(
                                new RunQueryFilter()
                                    .withOperand(RunQueryFilterOperand.PIPELINE_NAME)
                                    .withOperator(RunQueryFilterOperator.EQUALS)
                                    .withValues(Arrays.asList("examplePipeline")))),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.15/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.
