
/**
 * Samples for Replicas ListByConfigurationStore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresListReplicas.json
     */
    /**
     * Sample code: Replicas_ListByConfigurationStore.
     * 
     * @param manager Entry point to AppConfigurationManager.
     */
    public static void
        replicasListByConfigurationStore(com.azure.resourcemanager.appconfiguration.AppConfigurationManager manager) {
        manager.replicas().listByConfigurationStore("myResourceGroup", "contoso", null,
            com.azure.core.util.Context.NONE);
    }
}
