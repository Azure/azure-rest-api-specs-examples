Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-hdinsight_1.0.0-beta.5/sdk/hdinsight/azure-resourcemanager-hdinsight/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.hdinsight.models.Autoscale;
import com.azure.resourcemanager.hdinsight.models.AutoscaleConfigurationUpdateParameter;
import com.azure.resourcemanager.hdinsight.models.AutoscaleRecurrence;
import com.azure.resourcemanager.hdinsight.models.AutoscaleSchedule;
import com.azure.resourcemanager.hdinsight.models.AutoscaleTimeAndCapacity;
import com.azure.resourcemanager.hdinsight.models.DaysOfWeek;
import com.azure.resourcemanager.hdinsight.models.RoleName;
import java.util.Arrays;

/** Samples for Clusters UpdateAutoScaleConfiguration. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/EnableOrUpdateAutoScaleWithScheduleBasedConfiguration.json
     */
    /**
     * Sample code: Enable or Update Autoscale with the schedule based configuration for HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void enableOrUpdateAutoscaleWithTheScheduleBasedConfigurationForHDInsightCluster(
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
                            .withRecurrence(
                                new AutoscaleRecurrence()
                                    .withTimeZone("China Standard Time")
                                    .withSchedule(
                                        Arrays
                                            .asList(
                                                new AutoscaleSchedule()
                                                    .withDays(Arrays.asList(DaysOfWeek.THURSDAY))
                                                    .withTimeAndCapacity(
                                                        new AutoscaleTimeAndCapacity()
                                                            .withTime("16:00")
                                                            .withMinInstanceCount(4)
                                                            .withMaxInstanceCount(4)))))),
                Context.NONE);
    }
}
```
