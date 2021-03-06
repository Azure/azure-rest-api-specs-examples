import com.azure.core.util.Context;

/** Samples for Locations GetAzureAsyncOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Locations_GetAsyncOperationStatus.json
     */
    /**
     * Sample code: Gets the azure async operation status.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getsTheAzureAsyncOperationStatus(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .locations()
            .getAzureAsyncOperationStatusWithResponse(
                "East US 2", "8a0348f4-8a85-4ec2-abe0-03b26104a9a0-0", Context.NONE);
    }
}
