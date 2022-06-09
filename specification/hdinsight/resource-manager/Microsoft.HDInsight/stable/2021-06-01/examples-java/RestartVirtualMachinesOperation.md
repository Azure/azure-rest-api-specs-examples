```java
import com.azure.core.util.Context;
import java.util.Arrays;

/** Samples for VirtualMachines RestartHosts. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/RestartVirtualMachinesOperation.json
     */
    /**
     * Sample code: Restarts the specified HDInsight cluster hosts.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void restartsTheSpecifiedHDInsightClusterHosts(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.virtualMachines().restartHosts("rg1", "cluster1", Arrays.asList("gateway1", "gateway3"), Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.
