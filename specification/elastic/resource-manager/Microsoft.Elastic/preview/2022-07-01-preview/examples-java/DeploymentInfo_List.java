import com.azure.core.util.Context;

/** Samples for DeploymentInfo List. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2022-07-01-preview/examples/DeploymentInfo_List.json
     */
    /**
     * Sample code: DeploymentInfo_List.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void deploymentInfoList(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.deploymentInfoes().listWithResponse("myResourceGroup", "myMonitor", Context.NONE);
    }
}
