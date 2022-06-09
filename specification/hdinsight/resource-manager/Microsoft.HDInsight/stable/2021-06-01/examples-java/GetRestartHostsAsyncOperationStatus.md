```java
import com.azure.core.util.Context;

/** Samples for VirtualMachines GetAsyncOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetRestartHostsAsyncOperationStatus.json
     */
    /**
     * Sample code: Gets the async operation status of restarting host.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getsTheAsyncOperationStatusOfRestartingHost(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .virtualMachines()
            .getAsyncOperationStatusWithResponse(
                "rg1", "cluster1", "CF938302-6B4D-44A0-A6D2-C0D67E847AEC", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.
