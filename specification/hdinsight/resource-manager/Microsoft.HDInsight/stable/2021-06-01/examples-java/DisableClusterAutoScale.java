import com.azure.core.util.Context;
import com.azure.resourcemanager.hdinsight.models.AutoscaleConfigurationUpdateParameter;
import com.azure.resourcemanager.hdinsight.models.RoleName;

/** Samples for Clusters UpdateAutoScaleConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DisableClusterAutoScale.json
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
                "rg1", "cluster1", RoleName.WORKERNODE, new AutoscaleConfigurationUpdateParameter(), Context.NONE);
    }
}
