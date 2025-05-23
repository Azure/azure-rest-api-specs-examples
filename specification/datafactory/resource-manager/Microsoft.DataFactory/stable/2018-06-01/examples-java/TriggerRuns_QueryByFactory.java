
import com.azure.resourcemanager.datafactory.models.RunFilterParameters;
import com.azure.resourcemanager.datafactory.models.RunQueryFilter;
import com.azure.resourcemanager.datafactory.models.RunQueryFilterOperand;
import com.azure.resourcemanager.datafactory.models.RunQueryFilterOperator;
import java.time.OffsetDateTime;
import java.util.Arrays;

/**
 * Samples for TriggerRuns QueryByFactorySync.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * TriggerRuns_QueryByFactory.json
     */
    /**
     * Sample code: TriggerRuns_QueryByFactory.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void triggerRunsQueryByFactory(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.triggerRuns().queryByFactoryWithResponse("exampleResourceGroup", "exampleFactoryName",
            new RunFilterParameters().withLastUpdatedAfter(OffsetDateTime.parse("2018-06-16T00:36:44.3345758Z"))
                .withLastUpdatedBefore(OffsetDateTime.parse("2018-06-16T00:49:48.3686473Z"))
                .withFilters(Arrays.asList(new RunQueryFilter().withOperand(RunQueryFilterOperand.TRIGGER_NAME)
                    .withOperator(RunQueryFilterOperator.EQUALS).withValues(Arrays.asList("exampleTrigger")))),
            com.azure.core.util.Context.NONE);
    }
}
