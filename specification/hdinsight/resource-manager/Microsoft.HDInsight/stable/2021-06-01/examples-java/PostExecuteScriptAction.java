import com.azure.core.util.Context;
import com.azure.resourcemanager.hdinsight.models.ExecuteScriptActionParameters;
import com.azure.resourcemanager.hdinsight.models.RuntimeScriptAction;
import java.util.Arrays;

/** Samples for Clusters ExecuteScriptActions. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/PostExecuteScriptAction.json
     */
    /**
     * Sample code: Execute script action on HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void executeScriptActionOnHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .clusters()
            .executeScriptActions(
                "rg1",
                "cluster1",
                new ExecuteScriptActionParameters()
                    .withScriptActions(
                        Arrays
                            .asList(
                                new RuntimeScriptAction()
                                    .withName("Test")
                                    .withUri("http://testurl.com/install.ssh")
                                    .withParameters("")
                                    .withRoles(Arrays.asList("headnode", "workernode"))))
                    .withPersistOnSuccess(false),
                Context.NONE);
    }
}
