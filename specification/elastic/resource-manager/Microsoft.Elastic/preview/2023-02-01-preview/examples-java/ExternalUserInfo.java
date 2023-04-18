/** Samples for ExternalUser CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/preview/2023-02-01-preview/examples/ExternalUserInfo.json
     */
    /**
     * Sample code: ExternalUser_CreateOrUpdate.
     *
     * @param manager Entry point to ElasticManager.
     */
    public static void externalUserCreateOrUpdate(com.azure.resourcemanager.elastic.ElasticManager manager) {
        manager
            .externalUsers()
            .createOrUpdateWithResponse("myResourceGroup", "myMonitor", null, com.azure.core.util.Context.NONE);
    }
}
