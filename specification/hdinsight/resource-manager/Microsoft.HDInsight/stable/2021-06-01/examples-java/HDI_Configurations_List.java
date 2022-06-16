import com.azure.core.util.Context;

/** Samples for Configurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/HDI_Configurations_List.json
     */
    /**
     * Sample code: Get all configuration information.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAllConfigurationInformation(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.configurations().listWithResponse("rg1", "cluster1", Context.NONE);
    }
}
