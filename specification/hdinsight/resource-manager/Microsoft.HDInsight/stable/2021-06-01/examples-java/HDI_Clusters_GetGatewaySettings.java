import com.azure.core.util.Context;

/** Samples for Clusters GetGatewaySettings. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Clusters_GetGatewaySettings.json
     */
    /**
     * Sample code: Get HTTP settings.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getHTTPSettings(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.clusters().getGatewaySettingsWithResponse("rg1", "cluster1", Context.NONE);
    }
}
