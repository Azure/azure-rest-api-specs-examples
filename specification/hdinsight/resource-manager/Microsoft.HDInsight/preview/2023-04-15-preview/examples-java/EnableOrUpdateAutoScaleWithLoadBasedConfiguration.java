import com.azure.resourcemanager.hdinsight.models.Autoscale;
import com.azure.resourcemanager.hdinsight.models.AutoscaleCapacity;
import com.azure.resourcemanager.hdinsight.models.AutoscaleConfigurationUpdateParameter;
import com.azure.resourcemanager.hdinsight.models.RoleName;

/** Samples for Clusters UpdateAutoScaleConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/EnableOrUpdateAutoScaleWithLoadBasedConfiguration.json
     */
    /**
     * Sample code: Enable or Update Autoscale with the load based configuration for HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void enableOrUpdateAutoscaleWithTheLoadBasedConfigurationForHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .clusters()
            .updateAutoScaleConfiguration(
                "rg1",
                "cluster1",
                RoleName.WORKERNODE,
                new AutoscaleConfigurationUpdateParameter()
                    .withAutoscale(
                        new Autoscale()
                            .withCapacity(new AutoscaleCapacity().withMinInstanceCount(3).withMaxInstanceCount(5))),
                com.azure.core.util.Context.NONE);
    }
}
