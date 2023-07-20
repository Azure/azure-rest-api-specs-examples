import com.azure.resourcemanager.hdinsight.models.AutoscaleConfigurationUpdateParameter;
import com.azure.resourcemanager.hdinsight.models.RoleName;

/** Samples for Clusters UpdateAutoScaleConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/DisableClusterAutoScale.json
     */
    /**
     * Sample code: Disable Autoscale for the HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void disableAutoscaleForTheHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .clusters()
            .updateAutoScaleConfiguration(
                "rg1",
                "cluster1",
                RoleName.WORKERNODE,
                new AutoscaleConfigurationUpdateParameter(),
                com.azure.core.util.Context.NONE);
    }
}
