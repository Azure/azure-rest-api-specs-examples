Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-datafactory_1.0.0-beta.9/sdk/datafactory/azure-resourcemanager-datafactory/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.datafactory.models.DataFlowSink;
import com.azure.resourcemanager.datafactory.models.DataFlowSource;
import com.azure.resourcemanager.datafactory.models.DatasetReference;
import com.azure.resourcemanager.datafactory.models.MappingDataFlow;
import java.util.Arrays;

/** Samples for DataFlows CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/DataFlows_Create.json
     */
    /**
     * Sample code: DataFlows_Create.
     *
     * @param manager Entry point to DataFactoryManager.
     */
    public static void dataFlowsCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager
            .dataFlows()
            .define("exampleDataFlow")
            .withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withProperties(
                new MappingDataFlow()
                    .withDescription(
                        "Sample demo data flow to convert currencies showing usage of union, derive and conditional"
                            + " split transformation.")
                    .withSources(
                        Arrays
                            .asList(
                                new DataFlowSource()
                                    .withName("USDCurrency")
                                    .withDataset(new DatasetReference().withReferenceName("CurrencyDatasetUSD")),
                                new DataFlowSource()
                                    .withName("CADSource")
                                    .withDataset(new DatasetReference().withReferenceName("CurrencyDatasetCAD"))))
                    .withSinks(
                        Arrays
                            .asList(
                                new DataFlowSink()
                                    .withName("USDSink")
                                    .withDataset(new DatasetReference().withReferenceName("USDOutput")),
                                new DataFlowSink()
                                    .withName("CADSink")
                                    .withDataset(new DatasetReference().withReferenceName("CADOutput"))))
                    .withScript(
                        "source(output(PreviousConversionRate as double,Country as string,DateTime1 as"
                            + " string,CurrentConversionRate as double),allowSchemaDrift: false,validateSchema: false)"
                            + " ~> USDCurrency\n"
                            + "source(output(PreviousConversionRate as double,Country as string,DateTime1 as"
                            + " string,CurrentConversionRate as double),allowSchemaDrift: true,validateSchema: false)"
                            + " ~> CADSource\n"
                            + "USDCurrency, CADSource union(byName: true)~> Union\n"
                            + "Union derive(NewCurrencyRate = round(CurrentConversionRate*1.25)) ~> NewCurrencyColumn\n"
                            + "NewCurrencyColumn split(Country == 'USD',Country == 'CAD',disjoint: false) ~>"
                            + " ConditionalSplit1@(USD, CAD)\n"
                            + "ConditionalSplit1@USD sink(saveMode:'overwrite' ) ~> USDSink\n"
                            + "ConditionalSplit1@CAD sink(saveMode:'overwrite' ) ~> CADSink"))
            .create();
    }
}
```
