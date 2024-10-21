
/**
 * Samples for ExternalUser CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/ExternalUserInfo.json
     */
    /**
     * Sample code: ExternalUser_CreateOrUpdate.
     * 
     * @param manager Entry point to ElasticManager.
     */
    public static void externalUserCreateOrUpdate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager.externalUsers().createOrUpdateWithResponse("myResourceGroup", "myMonitor", null,
            com.azure.core.util.Context.NONE);
    }
}
