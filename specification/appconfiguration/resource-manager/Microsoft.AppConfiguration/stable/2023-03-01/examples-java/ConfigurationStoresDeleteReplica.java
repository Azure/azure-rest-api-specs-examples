/** Samples for Replicas Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2023-03-01/examples/ConfigurationStoresDeleteReplica.json
     */
    /**
     * Sample code: Replicas_Delete.
     *
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void replicasDelete(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.replicas().delete("myResourceGroup", "contoso", "myReplicaEus", com.azure.core.util.Context.NONE);
    }
}
