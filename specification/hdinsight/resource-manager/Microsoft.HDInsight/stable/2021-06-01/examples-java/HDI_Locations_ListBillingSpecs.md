```java
import com.azure.core.util.Context;

/** Samples for Locations ListBillingSpecs. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Locations_ListBillingSpecs.json
     */
    /**
     * Sample code: Get the subscription billingSpecs for the specified location.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getTheSubscriptionBillingSpecsForTheSpecifiedLocation(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.locations().listBillingSpecsWithResponse("East US 2", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.
