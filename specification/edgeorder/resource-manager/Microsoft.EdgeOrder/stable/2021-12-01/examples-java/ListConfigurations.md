Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-edgeorder_1.0.0-beta.1/sdk/edgeorder/azure-resourcemanager-edgeorder/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.edgeorder.models.ConfigurationFilters;
import com.azure.resourcemanager.edgeorder.models.ConfigurationsRequest;
import com.azure.resourcemanager.edgeorder.models.FilterableProperty;
import com.azure.resourcemanager.edgeorder.models.HierarchyInformation;
import com.azure.resourcemanager.edgeorder.models.SupportedFilterTypes;
import java.util.Arrays;

/** Samples for ResourceProvider ListConfigurations. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListConfigurations.json
     */
    /**
     * Sample code: ListConfigurations.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listConfigurations(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager
            .resourceProviders()
            .listConfigurations(
                new ConfigurationsRequest()
                    .withConfigurationFilters(
                        Arrays
                            .asList(
                                new ConfigurationFilters()
                                    .withHierarchyInformation(
                                        new HierarchyInformation()
                                            .withProductFamilyName("AzureStackEdge")
                                            .withProductLineName("AzureStackEdge")
                                            .withProductName("AzureStackEdgeGPU"))
                                    .withFilterableProperty(
                                        Arrays
                                            .asList(
                                                new FilterableProperty()
                                                    .withType(SupportedFilterTypes.SHIP_TO_COUNTRIES)
                                                    .withSupportedValues(Arrays.asList("US")))))),
                null,
                Context.NONE);
    }
}
```
