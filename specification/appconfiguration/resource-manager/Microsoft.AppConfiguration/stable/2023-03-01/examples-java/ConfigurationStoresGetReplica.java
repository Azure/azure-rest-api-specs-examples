/** Samples for Replicas Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresGetReplica.json
     */
    /**
     * Sample code: Replicas_Get.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void replicasGet(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager
            .replicas()
            .getWithResponse("myResourceGroup", "contoso", "myReplicaEus", com.azure.core.util.Context.NONE);
    }
}
