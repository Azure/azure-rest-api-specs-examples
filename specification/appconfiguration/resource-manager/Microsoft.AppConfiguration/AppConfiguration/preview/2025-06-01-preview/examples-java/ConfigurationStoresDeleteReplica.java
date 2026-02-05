
/**
 * Samples for Replicas Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresDeleteReplica.json
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
