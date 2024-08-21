
import com.azure.resourcemanager.hdinsight.models.UpdateGatewaySettingsParameters;

/**
 * Samples for Clusters UpdateGatewaySettings.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * HDI_Clusters_UpdateGatewaySettings_Enable.json
     */
    /**
     * Sample code: Enable HTTP connectivity.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void enableHTTPConnectivity(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.clusters().updateGatewaySettings("rg1", "cluster1", new UpdateGatewaySettingsParameters(),
            com.azure.core.util.Context.NONE);
    }
}
