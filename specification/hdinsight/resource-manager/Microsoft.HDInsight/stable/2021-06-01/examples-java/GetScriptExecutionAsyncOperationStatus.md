Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ScriptActions GetExecutionAsyncOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetScriptExecutionAsyncOperationStatus.json
     */
    /**
     * Sample code: Gets the async execution operation status.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getsTheAsyncExecutionOperationStatus(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .scriptActions()
            .getExecutionAsyncOperationStatusWithResponse(
                "rg1", "cluster1", "CF938302-6B4D-44A0-A6D2-C0D67E847AEC", Context.NONE);
    }
}
```
