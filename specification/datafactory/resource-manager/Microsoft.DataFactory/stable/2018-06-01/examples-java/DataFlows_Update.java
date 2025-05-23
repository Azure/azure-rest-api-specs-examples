
import com.azure.resourcemanager.datafactory.models.DataFlowResource;
import com.azure.resourcemanager.datafactory.models.DataFlowSink;
import com.azure.resourcemanager.datafactory.models.DataFlowSource;
import com.azure.resourcemanager.datafactory.models.DatasetReference;
import com.azure.resourcemanager.datafactory.models.MappingDataFlow;
import java.util.Arrays;

/**
 * Samples for DataFlows CreateOrUpdateSync.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Update.json
     */
    /**
     * Sample code: DataFlows_Update.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void dataFlowsUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        DataFlowResource resource = manager.dataFlows().getWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleDataFlow", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withProperties(new MappingDataFlow().withDescription(
            "Sample demo data flow to convert currencies showing usage of union, derive and conditional split transformation.")
            .withSources(Arrays.asList(
                new DataFlowSource().withName("USDCurrency")
                    .withDataset(new DatasetReference().withReferenceName("CurrencyDatasetUSD")),
                new DataFlowSource().withName("CADSource")
                    .withDataset(new DatasetReference().withReferenceName("CurrencyDatasetCAD"))))
            .withSinks(Arrays.asList(
                new DataFlowSink().withName("USDSink")
                    .withDataset(new DatasetReference().withReferenceName("USDOutput")),
                new DataFlowSink().withName("CADSink")
                    .withDataset(new DatasetReference().withReferenceName("CADOutput"))))
            .withScriptLines(Arrays.asList("source(output(", "PreviousConversionRate as double,", "Country as string,",
                "DateTime1 as string,", "CurrentConversionRate as double", "),", "allowSchemaDrift: false,",
                "validateSchema: false) ~> USDCurrency", "source(output(", "PreviousConversionRate as double,",
                "Country as string,", "DateTime1 as string,", "CurrentConversionRate as double", "),",
                "allowSchemaDrift: true,", "validateSchema: false) ~> CADSource",
                "USDCurrency, CADSource union(byName: true)~> Union",
                "Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn",
                "NewCurrencyColumn split(Country == 'USD',",
                "Country == 'CAD',disjoint: false) ~> ConditionalSplit1@(USD, CAD)",
                "ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink",
                "ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink")))
            .apply();
    }
}
