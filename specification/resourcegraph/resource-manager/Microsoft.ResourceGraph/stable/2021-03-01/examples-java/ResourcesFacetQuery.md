Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-resourcegraph_1.0.0-beta.3/sdk/resourcegraph/azure-resourcemanager-resourcegraph/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.resourcegraph.models.FacetRequest;
import com.azure.resourcemanager.resourcegraph.models.FacetRequestOptions;
import com.azure.resourcemanager.resourcegraph.models.FacetSortOrder;
import com.azure.resourcemanager.resourcegraph.models.QueryRequest;
import java.util.Arrays;

/** Samples for ResourceProvider Resources. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/stable/2021-03-01/examples/ResourcesFacetQuery.json
     */
    /**
     * Sample code: Query with a facet request.
     *
     * @param manager Entry point to ResourceGraphManager.
     */
    public static void queryWithAFacetRequest(com.azure.resourcemanager.resourcegraph.ResourceGraphManager manager) {
        manager
            .resourceProviders()
            .resourcesWithResponse(
                new QueryRequest()
                    .withSubscriptions(Arrays.asList("cfbbd179-59d2-4052-aa06-9270a38aa9d6"))
                    .withQuery(
                        "Resources | where type =~ 'Microsoft.Compute/virtualMachines' | project id, name, location,"
                            + " resourceGroup, properties.storageProfile.osDisk.osType | limit 5")
                    .withFacets(
                        Arrays
                            .asList(
                                new FacetRequest()
                                    .withExpression("location")
                                    .withOptions(
                                        new FacetRequestOptions().withSortOrder(FacetSortOrder.DESC).withTop(3)),
                                new FacetRequest()
                                    .withExpression("properties.storageProfile.osDisk.osType")
                                    .withOptions(
                                        new FacetRequestOptions().withSortOrder(FacetSortOrder.DESC).withTop(3)),
                                new FacetRequest()
                                    .withExpression("nonExistingColumn")
                                    .withOptions(
                                        new FacetRequestOptions().withSortOrder(FacetSortOrder.DESC).withTop(3)),
                                new FacetRequest()
                                    .withExpression("resourceGroup")
                                    .withOptions(
                                        new FacetRequestOptions()
                                            .withSortBy("tolower(resourceGroup)")
                                            .withSortOrder(FacetSortOrder.ASC)
                                            .withTop(3)),
                                new FacetRequest()
                                    .withExpression("resourceGroup")
                                    .withOptions(
                                        new FacetRequestOptions()
                                            .withFilter("resourceGroup contains 'test'")
                                            .withTop(3)))),
                Context.NONE);
    }
}
```
