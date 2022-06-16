import com.azure.core.util.Context;

/** Samples for Applications GetAzureAsyncOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetApplicationCreationAsyncOperationStatus.json
     */
    /**
     * Sample code: Get the azure async operation status.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getTheAzureAsyncOperationStatus(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .applications()
            .getAzureAsyncOperationStatusWithResponse(
                "rg1", "cluster1", "app", "CF938302-6B4D-44A0-A6D2-C0D67E847AEC", Context.NONE);
    }
}
