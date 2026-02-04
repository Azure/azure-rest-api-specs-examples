
/**
 * Samples for Replicas Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateReplica.json
     */
    /**
     * Sample code: Replicas_Create.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void replicasCreate(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.replicas().define("myReplicaEus").withExistingConfigurationStore("myResourceGroup", "contoso")
            .withRegion("eastus").create();
    }
}
